package storage

import (
	"github.com/reactivex/rxgo/handlers"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
	"github.com/sakari-ai/moirai/database"
	"github.com/sakari-ai/moirai/log"
	"github.com/sakari-ai/moirai/pkg/model"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type SQLStorage struct {
	DB database.DBEngine `inject:"database"`
}

const (
	CreateCell = "create"
	UpdateCell = "update"
)

type emitterKey struct {
	val      interface{}
	prop     model.PropertyType
	key      string
	recordID uuid.UUID
	id       uuid.UUID
	command  string
}

func (p SQLStorage) WriteSchema(schema *model.Schema) error {
	return p.DB.Model(&model.Schema{}).Create(schema).Error()
}

func (p SQLStorage) GetSchema(id uuid.UUID) (*model.Schema, error) {
	m := new(model.Schema)
	m.ID = id
	res := p.DB.Model(&model.Schema{}).Where("id = ?", id).
		First(m)

	return m, res.Error()
}

func (p SQLStorage) WriteRecord(schema model.Schema, record *model.Record) (*model.Record, error) {
	tx := p.DB.Begin()
	tx.Create(record)
	sequence := createEmissionFlat(schema, record)
	cellCtx := newCellEmission(tx)
	sequence = sequence.FlatMap(func(els interface{}) observable.Observable {
		return observable.Create(func(emitter *observer.Observer, disposed bool) {
			for _, item := range els.([]emitterKey) {
				emitter.OnNext(item)
			}
			emitter.OnDone()
		})
	}, 1)
	<-sequence.Subscribe(observer.New(cellCtx.next(), cellCtx.error(), cellCtx.done()))
	return record, nil
}

func (p SQLStorage) UpdateRecord(schema model.Schema, record *model.Record) (*model.Record, error) {
	tx := p.DB.Begin()
	exist := &model.Record{}
	rs := tx.Model(&model.Record{}).Find(exist, "id = ?", record.ID)

	if rs.Error() != nil {
		return exist, rs.Error()
	}

	record.ID = exist.ID
	sequence := updateEmissionFlat(schema, record)
	cellCtx := newCellEmission(tx)
	sequence = sequence.FlatMap(func(els interface{}) observable.Observable {
		return observable.Create(func(emitter *observer.Observer, disposed bool) {
			for _, item := range els.([]emitterKey) {
				emitter.OnNext(item)
			}
			emitter.OnDone()
		})
	}, 1)
	<-sequence.Subscribe(observer.New(cellCtx.next(), cellCtx.error(), cellCtx.done()))
	return exist, nil
}

func NewStorage() *SQLStorage {
	return new(SQLStorage)
}

type CellStorage interface {
	processCommand(prop emitterKey) error
}

func createRepo(prop model.PropertyType, tx database.DBEngine) CellStorage {
	switch prop.(type) {
	case *model.StringType:
		return createStrRepo(tx)
	case *model.FloatType:
		return createNumberRepo(tx)
	case *model.DateTimeType:
		return createDateRepo(tx)
	case *model.BooleanType:
		return createBoolRepo(tx)
	}
	return createIntRepo(tx)
}

func createEmissionFlat(schema model.Schema, record *model.Record) observable.Observable {
	emission := make([]emitterKey, 0, len(record.Fields.Columns))
	for k, v := range record.Fields.Columns {
		prop := schema.GetProp(k)
		if prop != nil {
			emission = append(emission, emitterKey{
				val:      v,
				prop:     prop,
				recordID: record.ID,
				command:  CreateCell,
				key:      k,
			})
		}
	}
	return observable.Just(emission)
}

func updateEmissionFlat(schema model.Schema, record *model.Record) observable.Observable {
	emission := make([]emitterKey, 0, len(record.Fields.Columns))
	for k, v := range record.Fields.Columns {
		prop := schema.GetProp(k)
		if prop != nil {
			emission = append(emission, emitterKey{
				val:      v,
				prop:     schema.Properties.Columns[k],
				recordID: record.ID,
				command:  UpdateCell,
				key:      k,
			})
		}
	}
	return observable.Just(emission)
}

type createEmitterContext struct {
	tx   database.DBEngine
	lock *sync.RWMutex
	errs bool
}

func newCellEmission(transaction database.DBEngine) *createEmitterContext {
	return &createEmitterContext{
		tx:   transaction,
		lock: new(sync.RWMutex),
	}
}

func (c *createEmitterContext) next() handlers.NextFunc {
	return func(el interface{}) {
		item := el.(emitterKey)
		c.lock.Lock()
		defer c.lock.Unlock()
		if !c.errs {
			cell := createRepo(item.prop, c.tx)
			if cell != nil {
				err := cell.processCommand(item)
				if err != nil {
					c.errs = true
				}
			}
		}
	}
}

func (c *createEmitterContext) error() handlers.ErrFunc {
	return func(err error) { log.Error("can not processCommand item") }
}

func (c *createEmitterContext) done() handlers.DoneFunc {
	return func() {
		if !c.errs {
			c.tx.Commit()
		} else {
			c.tx.Rollback()
		}
	}
}
