package handler

import (
	"github.com/sakari-ai/moirai/pkg/errors"
	"github.com/sakari-ai/moirai/proto"
	uuid "github.com/satori/go.uuid"
)

func validateRecords(records *proto.Records) []errors.FieldError {
	var errs []errors.FieldError
	_, err := uuid.FromString(records.ProjectID)
	if err != nil {
		errs = append(errs, errors.FieldError{Field: "project_id", Description: "invalid"})
	}
	_, err = uuid.FromString(records.SchemaID)
	if err != nil {
		errs = append(errs, errors.FieldError{Field: "schema_id", Description: "invalid"})
	}
	return errs
}

func validateRecord(record *proto.Record) []errors.FieldError {
	var errs []errors.FieldError
	_, err := uuid.FromString(record.ProjectID)
	if err != nil {
		errs = append(errs, errors.FieldError{Field: "project_id", Description: "invalid"})
	}
	_, err = uuid.FromString(record.SchemaID)
	if err != nil {
		errs = append(errs, errors.FieldError{Field: "schema_id", Description: "invalid"})
	}
	return errs
}
