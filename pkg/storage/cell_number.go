package storage

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/sakari-ai/moirai/database"
	uuid "github.com/satori/go.uuid"
)

type numberRepo struct {
	tx database.DBEngine
}

type NumberCell struct {
	ID       uuid.UUID `gorm:"column:id;primary_key"`
	RecordID uuid.UUID `gorm:"column:record_id"`
	Key      string    `gorm:"column:key"`
	Value    float64   `gorm:"column:value;"`
}

func (*NumberCell) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

func (i *numberRepo) processCommand(prop emitterKey) error {
	val, ok := prop.val.(float64)
	if !ok {
		return errors.New("value is not able to convert into float")
	}
	if prop.command == CreateCell {
		rs := i.tx.Create(&NumberCell{
			RecordID: prop.recordID,
			Key:      prop.key,
			Value:    val,
		})
		return rs.Error()
	}

	cell := &NumberCell{}
	res := i.tx.Model(&NumberCell{}).Find(cell, "record_id = ? and key = ?", prop.recordID, prop.key)
	if res.Error() != nil {
		return res.Error()
	}
	if cell.Value == val {
		return nil
	}

	cell.Value = val
	return i.tx.Save(cell).Error()
}

func createNumberRepo(tx database.DBEngine) *numberRepo {
	return &numberRepo{
		tx: tx,
	}
}
