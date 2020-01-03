package storage

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/sakari-ai/moirai/database"
	uuid "github.com/satori/go.uuid"
)

type stringRepo struct {
	tx database.DBEngine
}

type StringCell struct {
	ID       uuid.UUID `gorm:"column:id;primary_key"`
	RecordID uuid.UUID `gorm:"column:record_id"`
	Key      string    `gorm:"column:key"`
	Value    string    `gorm:"column:value;"`
}

func (*StringCell) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

func (s *stringRepo) processCommand(prop emitterKey) error {
	val, ok := prop.val.(string)
	if !ok {
		return errors.New("value is not able to convert into string")
	}
	if prop.command == CreateCell {
		return s.tx.Create(&StringCell{
			RecordID: prop.recordID,
			Key:      prop.key,
			Value:    val,
		}).Error()
	}

	cell := &StringCell{}
	res := s.tx.Model(&StringCell{}).Find(cell, "record_id = ? and key = ?", prop.recordID, prop.key)
	if res.Error() != nil {
		return res.Error()
	}
	if cell.Value == val {
		return nil
	}

	cell.Value = val
	return s.tx.Save(cell).Error()
}

func createStrRepo(tx database.DBEngine) *stringRepo {
	return &stringRepo{
		tx: tx,
	}
}
