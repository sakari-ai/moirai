package handler

import (
	"context"
	"fmt"
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
}

type Storage interface {
	WriteSchema(schema *model.Schema) error
	GetSchema(uuid uuid.UUID) (*model.Schema, error)
	WriteRecords(records []*model.Record) ([]*model.Record, []error)
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
	sch, _ := m.Storage.GetSchema(uuid.FromStringOrNil(req.Id))

	return transferSchemaToProto(sch), nil
}

func (m *Moirai) CreateRecords(ctx context.Context, records *proto.Records) (*proto.Records, error) {
	// Validate records

	// Convert proto -> model
	var mRecords []*model.Record
	for _, v := range records.Records {
		record, err := model.NewRecord(records.ProjectID, records.SchemaID, v.Fields)
		if err == nil {
			mRecords = append(mRecords, record)
		} else {
			fmt.Println("loi cai lon gi roi", err.Error())
		}
	}

	// Save it
	results, errs := m.WriteRecords(mRecords)

	// Process result & errors

	return transferRecordToProto(results), errors.BuildError(errs)
}

func (m *Moirai) UpdateRecords(context.Context, *proto.Records) (*proto.Records, error) {
	panic("implement me")
}
