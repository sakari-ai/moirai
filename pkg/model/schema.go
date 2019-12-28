package model

import (
	"database/sql/driver"
	"encoding/json"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Column struct {
	Key       string
	Type      string
	MinLength int32
	MaxLength int32
	Minimum   int64
	Maximum   int64
}

type PropertyType interface {
	Bind(p *DTOStruct) error
	ToProtoStruct() *structpb.Struct
}

type Properties struct {
	Columns map[string]PropertyType
}

type Schema struct {
	ID         uuid.UUID  `gorm:"column:id;primary_key"`
	Name       string     `gorm:"column:name"`
	Properties Properties `gorm:"column:properties;type:bytea"`
	Required   Required   `gorm:"column:required;type:bytea"`
	ProjectID  uuid.UUID  `gorm:"column:project_id"`

	Version   string `gorm:"column:version"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Required []string

func (c Required) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Required) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	return json.Unmarshal(v.([]byte), &c)
}

func (s *Schema) BeforeCreate(scope *gorm.Scope) error {
	ID := uuid.NewV4()
	err := scope.SetColumn("ID", ID)
	return err
}

func (c Properties) Value() (driver.Value, error) {
	return json.Marshal(c.Columns)
}

func (c *Properties) Scan(v interface{}) error {
	mapping := make(map[string]PropertyType)
	if v == nil {
		return nil
	}
	data := make(map[string]interface{})

	_ = json.Unmarshal(v.([]byte), &data)

	for k, val := range data {
		tpObject := struct {
			Type string `json:"type"`
		}{}
		rawJSON, _ := json.Marshal(val)

		_ = json.Unmarshal(rawJSON, &tpObject)
		var prop PropertyType
		switch tpObject.Type {
		case IntegerTp:
			prop = new(IntegerType)
		case FloatTp:
			prop = new(FloatType)
		case StringTp:
			prop = new(StringType)
		case BooleanTp:
			prop = new(BooleanType)
		case DateTp:
			prop = new(DateTimeType)
		}
		if prop != nil {
			err := json.Unmarshal(rawJSON, prop)
			if err == nil {
				mapping[k] = prop
			}
		}
	}
	c.Columns = mapping

	return nil
}
