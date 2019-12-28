package handler

import (
	"github.com/golang/protobuf/ptypes/struct"
	"github.com/sakari-ai/moirai/pkg/model"
	"github.com/sakari-ai/moirai/proto"
)

func transferSchemaToProto(m *model.Schema) *proto.Schema {
	sch := new(proto.Schema)

	sch.Required = m.Required
	sch.Name = m.Name
	sch.Id = m.ID.String()

	sch.Properties = map[string]*structpb.Struct{}

	for k, v := range m.Properties.Columns {
		sch.Properties[k] = v.ToProtoStruct()
	}
	sch.Version = m.Version

	return sch
}
