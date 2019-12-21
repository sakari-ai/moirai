package database

import (
	"github.com/jinzhu/gorm"
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

func ErrNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
