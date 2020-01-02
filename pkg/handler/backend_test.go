package handler

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/struct"
	"github.com/sakari-ai/moirai/core/util"
	"github.com/sakari-ai/moirai/pkg/errors"
	"github.com/sakari-ai/moirai/pkg/model"
	"github.com/sakari-ai/moirai/proto"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"time"
)

type mockStorage struct {
	mock.Mock
}

func (m *mockStorage) WriteSchema(schema *model.Schema) error {
	arg := m.Called(schema)

	return arg.Error(0)
}

func (m *mockStorage) GetSchema(uuid uuid.UUID) (*model.Schema, error) {
	arg := m.Called(uuid)

	return arg.Get(0).(*model.Schema), arg.Error(1)
}

func TestMoirai_CreateSchema(t *testing.T) {
	t.Parallel()
	type fields struct {
		Storage Storage
	}
	type args struct {
		ctx context.Context
		req *proto.Schema
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          *proto.Schema
		wantErr       bool
		expectedError error
	}{
		{
			name: "#1: Create schema success",
			fields: fields{Storage: func() *mockStorage {
				m := new(mockStorage)
				m.On("WriteSchema", mock.AnythingOfType("*model.Schema")).Return(nil)

				return m
			}()},
			args: args{
				ctx: context.Background(),
				req: &proto.Schema{
					Properties: map[string]*structpb.Struct{
						"name":   util.StructProto(map[string]interface{}{"type": "string"}),
						"age":    util.StructProto(map[string]interface{}{"type": "integer"}),
						"love":   util.StructProto(map[string]interface{}{"type": "boolean"}),
						"salary": util.StructProto(map[string]interface{}{"type": "number"}),
						"dob":    util.StructProto(map[string]interface{}{"type": "string", "format": "date-time"}),
					},
					Required:  []string{"name"},
					Name:      "NewSchema",
					ProjectID: uuid.NewV4().String(),
				},
			},
			want: &proto.Schema{
				Properties: map[string]*structpb.Struct{
					"name":   util.StructProto(map[string]interface{}{"type": "string"}),
					"age":    util.StructProto(map[string]interface{}{"type": "integer"}),
					"love":   util.StructProto(map[string]interface{}{"type": "boolean"}),
					"salary": util.StructProto(map[string]interface{}{"type": "number"}),
					"dob":    util.StructProto(map[string]interface{}{"type": "string", "format": "date-time"}),
				},
				Required: []string{"name"},
				Name:     "NewSchema",
			},
			wantErr: false,
		},
		{
			name: "#2: Create schema error",
			fields: fields{Storage: func() *mockStorage {
				m := new(mockStorage)
				m.On("WriteSchema", mock.AnythingOfType("*model.Schema")).Return(nil)

				return m
			}()},
			args: args{
				ctx: context.Background(),
				req: &proto.Schema{
					Properties: map[string]*structpb.Struct{
						"name": util.StructProto(map[string]interface{}{"type": "not-support"}),
					},
					Required:  []string{"name"},
					Name:      "NewSchema",
					ProjectID: uuid.NewV4().String(),
				},
			},
			want:    nil,
			wantErr: true,
			expectedError: errors.BuildInvalidArgument(errors.FieldError{
				Field:       "name",
				Description: "Field name has error (property (not-support) is not supported)",
			}),
		},
		{
			name: "#3: Create schema error",
			fields: fields{Storage: func() *mockStorage {
				m := new(mockStorage)
				m.On("WriteSchema", mock.AnythingOfType("*model.Schema")).Return(nil)

				return m
			}()},
			args: args{
				ctx: context.Background(),
				req: &proto.Schema{
					Properties: map[string]*structpb.Struct{
						"name": util.StructProto(map[string]interface{}{"type": "not-support"}),
					},
					Required:  []string{"name"},
					Name:      "NewSchema",
					ProjectID: "",
				},
			},
			want:    nil,
			wantErr: true,
			expectedError: errors.BuildInvalidArgument(errors.FieldError{
				Field:       "project_id",
				Description: "project_id is nil",
			}),
		},
		{
			name: "#4: Create schema db.error",
			fields: fields{Storage: func() *mockStorage {
				m := new(mockStorage)
				m.On("WriteSchema", mock.AnythingOfType("*model.Schema")).Return(fmt.Errorf("db error %s", "test 4"))

				return m
			}()},
			args: args{
				ctx: context.Background(),
				req: &proto.Schema{
					Properties: map[string]*structpb.Struct{
						"name": util.StructProto(map[string]interface{}{"type": "string"}),
					},
					Required:  []string{"name"},
					Name:      "NewSchema",
					ProjectID: uuid.NewV4().String(),
				},
			},
			want:          nil,
			wantErr:       true,
			expectedError: errors.Internal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Moirai{
				Storage: tt.fields.Storage,
			}
			got, err := m.CreateSchema(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				assert.Equal(t, err, tt.expectedError)
				return
			}
			assert.ObjectsAreEqualValues(got.Properties, tt.want.Properties)
			assert.ObjectsAreEqualValues(got.Required, tt.want.Required)
			assert.Equal(t, got.Name, tt.want.Name)
			assert.NotNil(t, got.Version, "Should have its own version")
		})
	}
}

func TestMoirai_GetSchema(t *testing.T) {
	type fields struct {
		Storage Storage
	}
	type args struct {
		ctx context.Context
		req *proto.RequestObjectById
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.Schema
		wantErr bool
	}{
		{

			name: "#1: Create schema success",
			fields: fields{Storage: func() *mockStorage {
				m := new(mockStorage)
				m.On("GetSchema", mock.Anything).Return(&model.Schema{
					ID:   uuid.NewV4(),
					Name: "FromDB",
					Properties: model.Properties{
						Columns: map[string]model.PropertyType{
							"name": &model.StringType{
								Type: "string",
							},
							"age": &model.IntegerType{
								Type: "integer",
							},
							"love": &model.BooleanType{
								Type: "boolean",
							},
							"salary": &model.FloatType{
								Type: "number",
							},
							"dob": &model.DateTimeType{
								Type:   "string",
								Format: "date-time",
							},
						},
					},
					Required:  model.Required{"name"},
					ProjectID: uuid.NewV4(),
					Version:   "xxxx",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)

				return m
			}()},
			args: args{
				ctx: context.Background(),
				req: &proto.RequestObjectById{
					Id: uuid.NewV4().String(),
				},
			},
			want: &proto.Schema{
				Properties: map[string]*structpb.Struct{
					"name":   util.StructProto(map[string]interface{}{"type": "string"}),
					"age":    util.StructProto(map[string]interface{}{"type": "integer"}),
					"love":   util.StructProto(map[string]interface{}{"type": "boolean"}),
					"salary": util.StructProto(map[string]interface{}{"type": "number"}),
					"dob":    util.StructProto(map[string]interface{}{"type": "string", "format": "date-time"}),
				},
				Required: []string{"name"},
				Name:     "FromDB",
				Version:  "xxxx",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Moirai{
				Storage: tt.fields.Storage,
			}
			got, err := m.GetSchema(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.ObjectsAreEqualValues(got.Properties, tt.want.Properties)
			assert.ObjectsAreEqualValues(got.Required, tt.want.Required)
			assert.ObjectsAreEqualValues(got.Name, tt.want.Name)
			assert.ObjectsAreEqualValues(got.Version, tt.want.Version)
		})
	}
}

func TestMoirai_Version(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *empty.Empty
	}
	tests := []struct {
		name    string
		args    args
		want    *proto.VersionResponse
		wantErr bool
	}{
		{
			name: "Get version",
			args: args{
				in0: context.Background(),
				in1: &empty.Empty{},
			},
			want: &proto.VersionResponse{
				Value: "unknown-version",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Moirai{}
			got, err := m.Version(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Version() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Version() got = %v, want %v", got, tt.want)
			}
		})
	}
}
