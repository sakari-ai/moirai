package handler

import (
	"context"
	errors2 "errors"
	"fmt"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/jinzhu/gorm"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	uuid "github.com/satori/go.uuid"

	"github.com/sakari-ai/moirai/log"
	"github.com/sakari-ai/moirai/log/field"
	"github.com/sakari-ai/moirai/pkg/errors"
	"github.com/sakari-ai/moirai/pkg/model"
	"github.com/sakari-ai/moirai/proto"
)

type Moirai struct {
	Storage `inject:"schema_storage"`

	Validator Validator `inject:"schema_validator"`
}

type Validator interface {
	Validate(schema model.Schema, record *structpb.Struct, opts ...model.JsonValidatorOption) []errors.FieldError
}

type Storage interface {
	WriteSchema(schema *model.Schema) error
	GetSchema(uuid uuid.UUID) (*model.Schema, error)
	WriteRecord(schema model.Schema, record *model.Record) (*model.Record, error)
	UpdateRecord(schema model.Schema, record *model.Record) (*model.Record, error)
}

func (m *Moirai) Version(context.Context, *empty.Empty) (*proto.VersionResponse, error) {
	version := os.Getenv("VERSION")
	if version == "" {
		version = "unknown-version"
	}
	return &proto.VersionResponse{Value: version}, nil
}

func (m *Moirai) CreateSchema(ctx context.Context, req *proto.Schema) (*proto.Schema, error) {
	projectID, err := uuid.FromString(req.ProjectID)
	if err != nil {
		return &proto.Schema{}, errors.BuildInvalidArgument(errors.FieldError{Field: "project_id", Description: "project_id is nil"})
	}
	schema, err := model.NewSchema(req.Name, projectID, req.Properties, req.Required...)
	if err != nil {
		return &proto.Schema{}, err
	}
	err = m.WriteSchema(schema)
	if err != nil {
		log.Error("write schema is failed", field.Error(err))
		return &proto.Schema{}, errors.Internal
	}
	req.Id = schema.ID.String()
	req.Version = schema.Version
	return req, nil
}

func (m *Moirai) GetSchema(ctx context.Context, req *proto.RequestObjectById) (*proto.Schema, error) {
	sch, err := m.Storage.GetSchema(uuid.FromStringOrNil(req.Id))
	if err != nil {
		if errors2.Is(err, gorm.ErrRecordNotFound) {
			return &proto.Schema{}, errors.BadError("item not found")
		}
	}

	return transferSchemaToProto(sch), nil
}

func (m *Moirai) CreateRecords(ctx context.Context, records *proto.Records) (*proto.Records, error) {
	sch, err := m.Storage.GetSchema(uuid.FromStringOrNil(records.SchemaID))
	if err != nil {
		if errors2.Is(err, gorm.ErrRecordNotFound) {
			return &proto.Records{}, errors.BadError("item not found")
		}
		log.Error("can not get schema", field.Error(err))
		return &proto.Records{}, errors.Internal
	}
	errs := validateRecords(records)
	if len(errs) > 0 {
		return nil, errors.BuildInvalidArgument(errs...)
	}
	pid, _ := uuid.FromString(records.ProjectID)
	sid, _ := uuid.FromString(records.SchemaID)

	var results []*model.Record
	for i, v := range records.Records {
		fErrors := m.Validator.Validate(*sch, v.Fields, model.WithFieldErrorPrefix(fmt.Sprintf("record.%d", i)))
		if len(fErrors) > 0 {
			errs = append(errs, fErrors...)
			continue
		}
		record := model.NewRecord(pid, sid, v.Fields)
		_, err = m.Storage.WriteRecord(*sch, record)
		if err != nil {
			log.Error("writing record has error", field.Any("record", record), field.Error(err))
			errs = append(errs, errors.FieldError{
				Field:       fmt.Sprintf("record.%d", i),
				Description: "Internal error",
			})
		}
		results = append(results, record)
	}
	if len(errs) > 0 {
		err = errors.BuildWithError("can not write schema", errs...)
	}

	return transferRecordToProto(results), err
}

func (m *Moirai) UpdateRecord(ctx context.Context, req *proto.Record) (*proto.Record, error) {
	sch, err := m.Storage.GetSchema(uuid.FromStringOrNil(req.SchemaID))
	if err != nil {
		if errors2.Is(err, gorm.ErrRecordNotFound) {
			return &proto.Record{}, errors.BadError("item not found")
		}
		log.Error("can not get schema", field.Error(err))
		return &proto.Record{}, errors.Internal
	}
	errs := validateRecord(req)
	if len(errs) > 0 {
		return nil, errors.BuildInvalidArgument(errs...)
	}
	pid, _ := uuid.FromString(req.ProjectID)
	sid, _ := uuid.FromString(req.SchemaID)

	fErrors := m.Validator.Validate(*sch, req.Fields)
	if len(fErrors) > 0 {
		errs = append(errs, fErrors...)
		return nil, errors.BuildInvalidArgument(errs...)
	}
	record := model.NewRecord(pid, sid, req.Fields)
	record.ID = uuid.FromStringOrNil(req.Id)
	_, err = m.Storage.UpdateRecord(*sch, record)
	if err != nil {
		if errors2.Is(err, gorm.ErrRecordNotFound) {
			return req, errors.NotFound
		}
		log.Error("writing req has error", field.Any("req", req), field.Error(err))
		return req, errors.BadError("can not update record")
	}
	return req, nil
}
