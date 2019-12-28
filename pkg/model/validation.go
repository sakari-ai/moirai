package model

import (
	"fmt"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/sakari-ai/moirai/core/util"
	"github.com/sakari-ai/moirai/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
	"sync"
)

type JsonSchemaValidator struct {
	schemaLoader map[string]*gojsonschema.Schema
	locker *sync.RWMutex
	loader *gojsonschema.SchemaLoader
}

func NewValidator()*JsonSchemaValidator {
	jLoader := &JsonSchemaValidator{
		schemaLoader: map[string]*gojsonschema.Schema{},
		locker:      new(sync.RWMutex),
		loader:       gojsonschema.NewSchemaLoader(),
	}
	jLoader.loader.Draft = gojsonschema.Draft7

	return jLoader
}

func (v *JsonSchemaValidator) Validate(schema Schema, record *structpb.Struct) (bool, error) {
	if _, ok := v.schemaLoader[schema.Version]; !ok {
		v.locker.Lock()
		defer v.locker.Unlock()
		sl := gojsonschema.NewStringLoader(schema.JSONSchema())
		sch, err := v.loader.Compile(sl)
		if err != nil {
			return false, err
		}
		v.schemaLoader[schema.Version] = sch
	}
	sch := v.schemaLoader[schema.Version]
	recordMap := util.StructToMap(record)
	documentLoader := gojsonschema.NewGoLoader(recordMap)
	result, err := sch.Validate(documentLoader)

	if err != nil {
		return false, err
	}
	var (
		errs    []errors.FieldError
		success = result.Valid()
	)
	for _, re := range result.Errors() {
		errs = append(errs, errors.FieldError{Field: re.Field(), Description: re.Description()})
	}
	if len(errs) > 0 {
		err = errors.BuildWithError(fmt.Sprintf("Record is not valid for Schema %s (Version: %s)",
			schema.Name,
			schema.Version),
			errs...)
	}
	return success, err
}
