package storage

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/sakari-ai/moirai/database"
	uuid "github.com/satori/go.uuid"
)

type boolRepo struct {
	tx database.DBEngine
}

type BoolCell struct {
	ID       uuid.UUID `gorm:"column:id;primary_key"`
	RecordID uuid.UUID `gorm:"column:record_id"`
	Key      string    `gorm:"column:key"`
	Value    bool      `gorm:"column:value;"`
}

func (*BoolCell) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

func (i *boolRepo) processCommand(prop emitterKey) error {
	val, ok := prop.val.(bool)
	if !ok {
		return errors.New("value is not able to convert into bool")
	}
	if prop.command == CreateCell {
		rs := i.tx.Create(&BoolCell{
			RecordID: prop.recordID,
			Key:      prop.key,
			Value:    val,
		})
		return rs.Error()
	}

	cell := &BoolCell{}
	res := i.tx.Model(&BoolCell{}).Find(cell, "record_id = ? and key = ?", prop.recordID, prop.key)
	if res.Error() != nil {
		return res.Error()
	}
	if cell.Value == val {
		return nil
	}

	cell.Value = val
	return i.tx.Save(cell).Error()
}

func createBoolRepo(tx database.DBEngine) *boolRepo {
	return &boolRepo{
		tx: tx,
	}
}
