package database

import (
	"github.com/jinzhu/gorm"
)

type Association interface {
	Find(value interface{}) Association
	Append(values ...interface{}) Association
	Replace(values ...interface{}) Association
	Delete(values ...interface{}) Association
	Count() int
	Error() error
}

type association struct {
	association *gorm.Association
}

func newAssociation(assn *gorm.Association) Association {
	return &association{association: assn}
}

func (a association) Find(value interface{}) Association {
	return newAssociation(a.association.Find(value))
}

func (a association) Append(values ...interface{}) Association {
	return newAssociation(a.association.Append(values...))
}

func (a association) Replace(values ...interface{}) Association {
	return newAssociation(a.association.Replace(values...))
}

func (a association) Delete(values ...interface{}) Association {
	return newAssociation(a.association.Delete(values...))
}

func (a association) Count() int {
	return a.association.Count()
}

func (a association) Error() error {
	return a.association.Error
}
