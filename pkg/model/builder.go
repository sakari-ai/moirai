package model

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/struct"
	"github.com/mitchellh/hashstructure"
	"github.com/sakari-ai/moirai/core/util"
	"github.com/sakari-ai/moirai/pkg/errors"
	"github.com/sakari-ai/moirai/proto"
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	DateTp    = "date-time"
	BooleanTp = "boolean"
	StringTp  = "string"
	IntegerTp = "integer"
	FloatTp   = "number"

	SchemeDefault = "default"
	SchemeMinimum = "minimum"
	SchemeMaximum = "maximum"
	TimeRFC3339   = "2006-01-02T15:04:05.999999999Z"
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
	if defVal := p.GetField(SchemeDefault); defVal != nil {
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
	if defVal := p.GetField(SchemeDefault); defVal != nil {
		defVal := int64(defVal.GetNumberValue())
		i.Default = defVal
	}
	if minVal := p.GetField(SchemeMinimum); minVal != nil {
		min := int64(minVal.GetNumberValue())
		i.Minimum = min
	}
	if maxVal := p.GetField(SchemeMaximum); maxVal != nil {
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
	if defVal := p.GetField(SchemeDefault); defVal != nil {
		defVal := defVal.GetNumberValue()
		fl.Default = defVal
	}
	if minVal := p.GetField(SchemeMinimum); minVal != nil {
		min := minVal.GetNumberValue()
		fl.Minimum = min
	}
	if maxVal := p.GetField(SchemeMaximum); maxVal != nil {
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

func (b *DateTimeType) ToProtoStruct() *structpb.Struct {
	return bindingStruct(b)
}

func (b *DateTimeType) Bind(p *DTOStruct) error {
	if desVal := p.GetField("description"); desVal != nil {
		b.Description = desVal.GetStringValue()
	}
	if defVal := p.GetField(SchemeDefault); defVal != nil {
		_, err := time.Parse(time.RFC3339, defVal.GetStringValue())
		if err != nil {
			return err
		}
		b.Default = defVal.GetStringValue()
	}
	b.Type = DateTp

	return nil
}

func (b *DateTimeType) Load() {
	b.Type = StringTp
	b.Format = "date-time"
}

func CreateProperty(p *DTOStruct) (PropertyType, error) {
	tp := p.GetField("type")
	format := p.GetField("format")
	if tp == nil {
		return nil, errors.BadError("type not found")
	}
	if tp.GetStringValue() == StringTp && format.GetStringValue() == "" {
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
	if tp.GetStringValue() == StringTp && format.GetStringValue() == DateTp {
		prop := new(DateTimeType)
		err := prop.Bind(p)
		prop.Load()
		return prop, err
	}
	return nil, fmt.Errorf("property (%s) is not supported", tp.GetStringValue())
}

func NewSchema(name string, projectID uuid.UUID, columns map[string]*structpb.Struct, required ...string) (*Schema, error) {
	sch := &Schema{Name: name, Required: required}
	var errs []errors.FieldError
	sch.ProjectID = projectID
	properties := Properties{Columns: map[string]PropertyType{}}
	for k, val := range columns {
		column := DTOStruct(*val)
		prop, err := CreateProperty(&column)
		if err != nil {
			errs = append(errs, errors.FieldError{Field: k, Description: fmt.Sprintf("Field %s has error (%s)", k, err.Error())})
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
