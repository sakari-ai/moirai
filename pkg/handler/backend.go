package handler

import (
	"context"
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
}

func (m *Moirai) Version(context.Context, *empty.Empty) (*proto.VersionResponse, error) {
	version := os.Getenv("VERSION")
	if version == "" {
		version = "unknown-version"
	}
	return &proto.VersionResponse{Value: version}, nil
}

func (m *Moirai) CreateSchema(ctx context.Context, req *proto.Schema) (*proto.Schema, error) {
	schema, err := model.NewSchema(req.Name, req.ProjectID, req.Properties)
	if err != nil {
		return &proto.Schema{}, err
	}
	err = m.WriteSchema(schema)
	if err != nil {
		log.Error("write schema is failed", field.Error(err))
		return &proto.Schema{}, errors.Internal
	}
	req.Id = schema.ID.String()
	return req, nil
}

func (m *Moirai) GetSchema(ctx context.Context, req *proto.RequestObjectById) (*proto.Schema, error) {
	model, _ := m.Storage.GetSchema(uuid.FromStringOrNil(req.Id))

	log.Info("schema", field.Any("schema", model))
	return &proto.Schema{}, nil
}
