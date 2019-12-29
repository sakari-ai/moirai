package handler

import (
	"github.com/golang/protobuf/ptypes/struct"
	"github.com/sakari-ai/moirai/core/util"
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

func transferRecordToProto(m []*model.Record) *proto.Records {
	if len(m) < 1 {
		return nil
	}
	records := new(proto.Records)
	records.SchemaID = m[0].SchemaID.String()
	records.ProjectID = m[0].ProjectID.String()

	for _, v := range m {
		record := new(proto.Record)
		record.Id = v.ID.String()
		record.Fields = util.StructProto(v.Fields.Columns)
		records.Records = append(records.Records, record)
	}

	return records
}
