package handler

import (
	"context"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sakari-ai/moirai/proto"
)

type Moirai struct {
}

func (m *Moirai) Version(context.Context, *empty.Empty) (*proto.VersionResponse, error) {
	version := os.Getenv("VERSION")
	if version == "" {
		version = "unknown-version"
	}
	return &proto.VersionResponse{Value: version}, nil
}

func (m *Moirai) CreateSchema(ctx context.Context, req *proto.Schema) (*proto.Schema, error) {
	return &proto.Schema{}, nil
}
