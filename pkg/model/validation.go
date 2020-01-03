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
	locker       *sync.RWMutex
	loader       *gojsonschema.SchemaLoader
}

func NewValidator() *JsonSchemaValidator {
	jLoader := &JsonSchemaValidator{
		schemaLoader: map[string]*gojsonschema.Schema{},
		locker:       new(sync.RWMutex),
		loader:       gojsonschema.NewSchemaLoader(),
	}
	jLoader.loader.Draft = gojsonschema.Draft7

	return jLoader
}

type processContext struct {
	errorHandler func(*errors.FieldError)
}

type JsonValidatorOption func(*processContext)

func WithFieldErrorPrefix(prefix string) JsonValidatorOption {
	return func(context *processContext) {
		context.errorHandler = func(fieldError *errors.FieldError) {
			fieldError.Field = fmt.Sprintf("%s%s", prefix, fieldError.Field)
		}
	}
}

func (v *JsonSchemaValidator) Validate(schema Schema, record *structpb.Struct, opts ...JsonValidatorOption) []errors.FieldError {
	ctx := &processContext{errorHandler: func(fieldError *errors.FieldError) {}}
	for _, f := range opts {
		f(ctx)
	}
	var (
		errs []errors.FieldError
	)
	if _, ok := v.schemaLoader[schema.Version]; !ok {
		v.locker.Lock()
		defer v.locker.Unlock()
		sl := gojsonschema.NewStringLoader(schema.JSONSchema())
		sch, err := v.loader.Compile(sl)
		if err != nil {
			errs = append(errs, errors.FieldError{
				Field:       "schema definition has error",
				Description: err.Error(),
			})
		}
		v.schemaLoader[schema.Version] = sch
	}
	sch := v.schemaLoader[schema.Version]
	recordMap := util.StructToMap(record)
	documentLoader := gojsonschema.NewGoLoader(recordMap)
	result, err := sch.Validate(documentLoader)
	if err != nil {
		errs = append(errs, errors.FieldError{
			Field:       "document schema has error",
			Description: err.Error(),
		})
	}
	if result != nil {
		for _, re := range result.Errors() {
			fe := errors.FieldError{Field: re.Field(), Description: re.Description()}
			ctx.errorHandler(&fe)
			errs = append(errs, fe)
		}
		if len(errs) > 0 {
			err = errors.BuildWithError(fmt.Sprintf("Record is not valid for Schema %s (Version: %s)",
				schema.Name,
				schema.Version),
				errs...)
		}
	}
	return errs
}
