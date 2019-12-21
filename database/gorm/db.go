package gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	postgres = "postgres"
)

type DB = gorm.DB

type DSN string

func (d DSN) String() string {
	return string(d)
}

func Open(dsn DSN) (*DB, error) {
	db, err := gorm.Open(postgres, dsn.String())
	return db, err
}

func OpenDialects(dialect string, dsn DSN) (*DB, error) {
	db, err := gorm.Open(dialect, dsn.String())
	return db, err
}
