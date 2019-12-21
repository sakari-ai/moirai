package bootstrap

import (
	"github.com/sakari-ai/moirai/config"
	"github.com/sakari-ai/moirai/database"
	"github.com/sakari-ai/moirai/database/gorm"
	"github.com/sakari-ai/moirai/migration"
)

var dbEngine database.DBEngine

func LoadDB(cfg *config.Database) {
	if cfg == nil {
		panic("config: nil")
	}
	dbcfg := database.Config(*cfg)
	err := runMigration(dbcfg)
	if err != nil {
		panic(err)
	}

	dbEngine, err = database.Open(dbcfg)
	if dbcfg.Debug {
		dbEngine.Debug()
		dbEngine.LogMode(true)
	}
	if err != nil {
		panic(err)
	}
}

func runMigration(cfg database.Config) error {
	db, err := gorm.Open(cfg.DSN())
	db.LogMode(true)
	if err != nil {
		return nil
	}
	defer db.Close()
	m := migration.New(migration.WithGormigrate(db, cfg.Migrate))
	return m.Migrate()
}
