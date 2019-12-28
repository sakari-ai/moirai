package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/sakari-ai/moirai/pkg/model"
	"gopkg.in/gormigrate.v1"
)

func CoreMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "schema-database",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(&model.Schema{}).Error
			},
		},
	}
}
