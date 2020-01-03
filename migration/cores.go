package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/sakari-ai/moirai/pkg/model"
	"github.com/sakari-ai/moirai/pkg/storage"
	"gopkg.in/gormigrate.v1"
)

func CoreMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "schema-database",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(
					&model.Schema{},
				).Error
			},
		},
		{
			ID: "add-record",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(
					&model.Record{},
				).Error
			},
		},
		{
			ID: "add-record-cells",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(
					&storage.BoolCell{},
					&storage.IntCell{},
					&storage.StringCell{},
					&storage.DateCell{},
					&storage.BoolCell{},
					&storage.NumberCell{},
				).Error
			},
		},
	}
}
