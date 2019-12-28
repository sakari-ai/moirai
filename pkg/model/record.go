package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Fields struct {
	Columns map[string]PropertyType
}

type Record struct {
	ID        uuid.UUID `gorm:"column:id;primary_key"`
	ProjectID uuid.UUID `gorm:"column:project_id"`
	SchemaID  uuid.UUID `gorm:"column:schema_id"`
	Fields    Fields    `gorm:"column:fields;type:bytea"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
