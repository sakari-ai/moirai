package model

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/struct"
	"github.com/mitchellh/hashstructure"
	"github.com/sakari-ai/moirai/core/util"
	"github.com/sakari-ai/moirai/internal/lib/ptypes"
	"github.com/sakari-ai/moirai/pkg/errors"
	"github.com/sakari-ai/moirai/proto"
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	DateTp    = "date"
	BooleanTp = "boolean"
	StringTp  = "string"
	IntegerTp = "integer"
	FloatTp   = "float"
)

type DTOStruct structpb.Struct

type Columns map[string]*structpb.Struct

func (d *DTOStruct) GetField(field string) *structpb.Value {
	return d.Fields[field]
}

type StringType proto.StringType

func bindingStruct(p interface{}) *structpb.Struct {
	raw := make(map[string]interface{})
	bts, _ := json.Marshal(p)
	json.Unmarshal(bts, &raw)

	return util.StructProto(raw)
}

func (s *StringType) ToProtoStruct() *structpb.Struct {
	return bindingStruct(s)
}

func (s *StringType) Bind(p *DTOStruct) error {
	if defVal := p.GetField("default"); defVal != nil {
		s.Default = defVal.GetStringValue()
	}
	if minVal := p.GetField("minLength"); minVal != nil {
		s.MinLength = int32(minVal.GetNumberValue())
	}
	if maxLen := p.GetField("maxLength"); maxLen != nil {
		s.MaxLength = int32(maxLen.GetNumberValue())
	}
	if desVal := p.GetField("description"); desVal != nil {
		s.Description = desVal.GetStringValue()
	}
	s.Type = StringTp

	return nil
}

type IntegerType proto.IntegerType

func (i *IntegerType) ToProtoStruct() *structpb.Struct {
	return bindingStruct(i)
}

func (i *IntegerType) Bind(p *DTOStruct) error {
	if defVal := p.GetField("default"); defVal != nil {
		defVal := int64(defVal.GetNumberValue())
		i.Default = defVal
	}
	if minVal := p.GetField("minimum"); minVal != nil {
		min := int64(minVal.GetNumberValue())
		i.Minimum = min
	}
	if maxVal := p.GetField("maximum"); maxVal != nil {
		max := int64(maxVal.GetNumberValue())
		i.Maximum = max
	}
	if desVal := p.GetField("description"); desVal != nil {
		i.Description = desVal.GetStringValue()
	}
	i.Type = IntegerTp

	return nil
}

type FloatType proto.FloatType

func (fl *FloatType) ToProtoStruct() *structpb.Struct {
	return bindingStruct(fl)
}

func (fl *FloatType) Bind(p *DTOStruct) error {
	if defVal := p.GetField("default"); defVal != nil {
		defVal := defVal.GetNumberValue()
		fl.Default = defVal
	}
	if minVal := p.GetField("minimum"); minVal != nil {
		min := minVal.GetNumberValue()
		fl.Minimum = min
	}
	if maxVal := p.GetField("maximum"); maxVal != nil {
		max := maxVal.GetNumberValue()
		fl.Maximum = max
	}
	if desVal := p.GetField("description"); desVal != nil {
		fl.Description = desVal.GetStringValue()
	}
	fl.Type = FloatTp

	return nil
}

type BooleanType proto.BooleanType

func (b *BooleanType) ToProtoStruct() *structpb.Struct {
	return bindingStruct(b)
}

func (b *BooleanType) Bind(p *DTOStruct) error {
	if defVal := p.GetField("description"); defVal != nil {
		b.Description = defVal.GetStringValue()
	}
	b.Type = BooleanTp

	return nil
}

type DateTimeType proto.DateTimeType

func (d *DateTimeType) ToProtoStruct() *structpb.Struct {
	return bindingStruct(d)
}

func (b *DateTimeType) Bind(p *DTOStruct) error {
	if desVal := p.GetField("description"); desVal != nil {
		b.Description = desVal.GetStringValue()
	}

	if minVal := p.GetField("minimum"); minVal != nil {
		min, err := time.Parse(time.RFC3339, minVal.GetStringValue())
		if err != nil {
			return err
		}
		b.Minimum = ptypes.TimestampProto(min)
	}
	if maxVal := p.GetField("maximum"); maxVal != nil {
		max, err := time.Parse(time.RFC3339, p.Fields["maximum"].GetStringValue())
		if err != nil {
			return err
		}
		b.Maximum = ptypes.TimestampProto(max)
	}
	if defVal := p.GetField("default"); defVal != nil {
		def, err := time.Parse(time.RFC3339, p.Fields["default"].GetStringValue())
		if err != nil {
			return err
		}
		b.Default = ptypes.TimestampProto(def)
	}
	b.Type = DateTp

	return nil
}

func CreateProperty(p *DTOStruct) (PropertyType, error) {
	tp := p.GetField("type")
	if tp == nil {
		return nil, errors.BadError("type not found")
	}
	if tp.GetStringValue() == StringTp {
		prop := new(StringType)
		err := prop.Bind(p)
		return prop, err
	}
	if tp.GetStringValue() == IntegerTp {
		prop := new(IntegerType)
		err := prop.Bind(p)
		return prop, err
	}
	if tp.GetStringValue() == FloatTp {
		prop := new(FloatType)
		err := prop.Bind(p)
		return prop, err
	}
	if tp.GetStringValue() == BooleanTp {
		prop := new(BooleanType)
		err := prop.Bind(p)
		return prop, err
	}
	if tp.GetStringValue() == DateTp {
		prop := new(DateTimeType)
		err := prop.Bind(p)
		return prop, err
	}

	return nil, errors.BadError("property not found")
}

func NewSchema(name, projectId string, columns map[string]*structpb.Struct, required ...string) (*Schema, error) {
	sch := &Schema{Name: name, Required: required}
	var errs []errors.FieldError
	projectID, err := uuid.FromString(projectId)
	if err != nil {
		errs = append(errs, errors.FieldError{Field: "project_id", Description: "Not nill"})
		return sch, errors.BuildInvalidArgument(errs...)
	}
	sch.ProjectID = projectID
	properties := Properties{Columns: map[string]PropertyType{}}
	for k, val := range columns {
		column := DTOStruct(*val)
		prop, err := CreateProperty(&column)
		if err != nil {
			errs = append(errs, errors.FieldError{Field: k, Description: fmt.Sprintf("Field %s has error %s", k, err.Error())})
		} else {
			properties.Columns[k] = prop
		}
	}
	if len(errs) == 0 {
		sch.Properties = properties
		version, _ := hashstructure.Hash(sch, nil)
		sch.Version = fmt.Sprint(version)
		return sch, nil
	}
	return sch, errors.BuildInvalidArgument(errs...)
}
