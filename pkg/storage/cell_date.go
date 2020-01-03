package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sakari-ai/moirai/database"
	"github.com/sakari-ai/moirai/pkg/model"
	uuid "github.com/satori/go.uuid"
	"time"
)

type dateRepo struct {
	tx database.DBEngine
}

type DateCell struct {
	ID       uuid.UUID `gorm:"column:id;primary_key"`
	RecordID uuid.UUID `gorm:"column:record_id"`
	Key      string    `gorm:"column:key"`
	Value    time.Time `gorm:"column:value"`
}

func (*DateCell) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

func (d *dateRepo) processCommand(prop emitterKey) error {
	dt, err := time.Parse(model.TimeRFC3339, fmt.Sprint(prop.val))
	if err != nil {
		return err
	}
	if prop.command == CreateCell {
		return d.tx.Create(&DateCell{
			RecordID: prop.recordID,
			Key:      prop.key,
			Value:    dt,
		}).Error()
	}

	cell := &DateCell{}
	res := d.tx.Model(&DateCell{}).Find(cell, "record_id = ? and key = ?", prop.recordID, prop.key)
	if res.Error() != nil {
		return res.Error()
	}
	if cell.Value == dt {
		return nil
	}

	cell.Value = dt
	return d.tx.Save(cell).Error()
}

func createDateRepo(tx database.DBEngine) *dateRepo {
	return &dateRepo{
		tx: tx,
	}
}
