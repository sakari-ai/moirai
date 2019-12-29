package storage

import (
	"github.com/sakari-ai/moirai/database"
	"github.com/sakari-ai/moirai/pkg/model"
	uuid "github.com/satori/go.uuid"
)

type PostgresStorage struct {
	DB database.DBEngine `inject:"database"`
}

func (p PostgresStorage) WriteSchema(schema *model.Schema) error {
	return p.DB.Model(&model.Schema{}).Create(schema).Error()
}

func (p PostgresStorage) GetSchema(id uuid.UUID) (*model.Schema, error) {
	m := new(model.Schema)
	m.ID = id
	res := p.DB.Model(&model.Schema{}).Where("id = ?", id).
		First(m)

	return m, res.Error()
}

func (p PostgresStorage) WriteRecords(records []*model.Record) ([]*model.Record, []error) {
	var errs []error
	for _, record := range records {
		if err := p.DB.Create(&record).Error(); err != nil {
			errs = append(errs, err)
		}
	}
	return records, errs
}

func NewStorage() *PostgresStorage {
	return new(PostgresStorage)
}
