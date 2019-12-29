package model

import (
	"database/sql/driver"
	"encoding/json"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/jinzhu/gorm"
	"github.com/sakari-ai/moirai/core/util"
	"github.com/sakari-ai/moirai/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Fields struct {
	Columns map[string]interface{}
}

type Record struct {
	ID        uuid.UUID `gorm:"column:id;primary_key"`
	ProjectID uuid.UUID `gorm:"column:project_id"`
	SchemaID  uuid.UUID `gorm:"column:schema_id"`
	Fields    Fields    `gorm:"column:fields;type:bytea"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *Record) BeforeCreate(scope *gorm.Scope) error {
	ID := uuid.NewV4()
	err := scope.SetColumn("ID", ID)
	return err
}

func (c Fields) Value() (driver.Value, error) {
	return json.Marshal(c.Columns)
}

func (c *Fields) Scan(v interface{}) error {
	if v == nil {
		return nil
	}
	data := make(map[string]interface{})

	_ = json.Unmarshal(v.([]byte), &data)

	c.Columns = data

	return nil
}

func NewRecord(projectID string, schemaID string, fields *structpb.Struct) (*Record, error) {
	record := &Record{}
	var errs []errors.FieldError
	pid, err := uuid.FromString(projectID)
	if err != nil {
		errs = append(errs, errors.FieldError{Field: "project_id", Description: "invalid"})
	}

	sid, err := uuid.FromString(schemaID)
	if err != nil {
		errs = append(errs, errors.FieldError{Field: "schema_id", Description: "invalid"})
	}

	if len(errs) > 0 {
		return nil, errors.BuildInvalidArgument(errs...)
	}

	record.ProjectID = pid
	record.SchemaID = sid
	record.Fields = Fields{
		Columns: util.StructToMap(fields),
	}

	return record, nil
}
