package storage

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sakari-ai/moirai/database"
	uuid "github.com/satori/go.uuid"
)

type intRepo struct {
	tx database.DBEngine
}

type IntCell struct {
	ID       uuid.UUID `gorm:"column:id;primary_key"`
	RecordID uuid.UUID `gorm:"column:record_id"`
	Key      string    `gorm:"column:key"`
	Value    int64     `gorm:"column:value;type:integer"`
}

func (*IntCell) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

func (i *intRepo) processCommand(prop emitterKey) error {
	val, ok := prop.val.(int)
	if !ok {
		valF, ok := prop.val.(float64)
		if ok {
			val = int(valF)
		}
		if !ok {
			return errors.New("value is not able to convert into int")
		}
	}
	if prop.command == CreateCell {
		return i.tx.Create(&IntCell{
			RecordID: prop.recordID,
			Key:      prop.key,
			Value:    int64(val),
		}).Error()
	}

	cell := &IntCell{}
	res := i.tx.Model(&IntCell{}).Find(cell, "record_id = ? and key = ?", prop.recordID, prop.key)
	if res.Error() != nil {
		return res.Error()
	}
	if cell.Value == int64(val) {
		return nil
	}

	cell.Value = int64(val)
	return i.tx.Save(cell).Error()
}

func createIntRepo(tx database.DBEngine) *intRepo {
	return &intRepo{
		tx: tx,
	}
}
