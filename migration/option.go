package migration

import (
	"github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

type Option func(m *Migration)

var (
	preset = map[string][]*gormigrate.Migration{}
)

func WithGormigrate(db *gorm.DB, mod string) Option {
	return func(m *Migration) {
		var migrations []*gormigrate.Migration
		if migration, ok := preset[mod]; ok {
			migrations = append(migrations, migration...)
			mr := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
			m.mr = mr
		}
	}
}
