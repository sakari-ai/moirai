package handler

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/struct"
	"github.com/sakari-ai/moirai/core/util"
	"github.com/sakari-ai/moirai/database"
	"github.com/sakari-ai/moirai/database/fake"
	"github.com/sakari-ai/moirai/pkg/errors"
	"github.com/sakari-ai/moirai/pkg/model"
	"github.com/sakari-ai/moirai/pkg/storage"
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

func (m *mockStorage) WriteRecord(schema model.Schema, record *model.Record) (*model.Record, error) {
	arg := m.Called(schema, record)

	return arg.Get(0).(*model.Record), arg.Error(1)
}

func (m *mockStorage) UpdateRecord(schema model.Schema, record *model.Record) (*model.Record, error) {
	arg := m.Called(schema, record)

	return arg.Get(0).(*model.Record), arg.Error(1)
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

func TestMoirai_UpdateRecord(t *testing.T) {
	sch := &model.Schema{
		Name: "Simple One",
		Properties: model.Properties{
			Columns: map[string]model.PropertyType{
				"name": &model.StringType{
					Type:        model.StringTp,
					Description: "description str",
					MinLength:   1,
					MaxLength:   10,
					Default:     "paul",
				},
				"age":    &model.IntegerType{Type: model.IntegerTp},
				"salary": &model.FloatType{Type: model.FloatTp},
				"love":   &model.BooleanType{Type: model.BooleanTp},
				"dob":    &model.DateTimeType{Type: model.StringTp, Format: model.DateTp},
			},
		},
		Required:  []string{},
		ProjectID: uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57c"),
	}
	rcd := &model.Record{
		ProjectID: uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57c"),
		SchemaID:  sch.ID,
		Fields: model.Fields{
			Columns: map[string]interface{}{
				"name":   "paul",
				"age":    35,
				"salary": 3.5,
				"love":   true,
				"dob":    time.Date(1986, 12, 28, 0, 0, 0, 0, time.Local).Format(model.TimeRFC3339),
			},
		},
	}
	type fields struct {
		DB      database.DBEngine
		Storage Storage
	}
	type args struct {
		schema model.Schema
		record *proto.Record
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Record
		wantErr bool
	}{
		{
			name: "#1:  Update Records simple",
			fields: fields{
				DB: func() database.DBEngine {
					mockDB := fake.PrepareForTesting(t, fake.WithModels(&model.Schema{}, &model.Record{}, &storage.IntCell{}, &storage.NumberCell{}, &storage.StringCell{}, &storage.DateCell{}, &storage.BoolCell{}))

					mockDB.Create(sch)
					mockDB.Create(rcd)
					mockDB.Create(&storage.StringCell{
						RecordID: rcd.ID,
						Key:      "name",
						Value:    "paul",
					})
					mockDB.Create(&storage.IntCell{
						RecordID: rcd.ID,
						Key:      "age",
						Value:    35,
					})
					mockDB.Create(&storage.NumberCell{
						RecordID: rcd.ID,
						Key:      "salary",
						Value:    3.5,
					})
					mockDB.Create(&storage.BoolCell{
						RecordID: rcd.ID,
						Key:      "love",
						Value:    true,
					})
					mockDB.Create(&storage.DateCell{
						ID:       uuid.NewV4(),
						RecordID: rcd.ID,
						Key:      "dob",
						Value:    time.Date(1986, 12, 28, 0, 0, 0, 0, time.Local),
					})
					return mockDB
				}(),
			},
			args: args{
				schema: *sch,
				record: &proto.Record{
					Id:        rcd.ID.String(),
					ProjectID: "3341fa1e-90b0-482a-b0ac-74a76d6af57c",
					SchemaID:  sch.ID.String(),
					Fields: util.StructProto(map[string]interface{}{
						"name":   "dima",
						"age":    31,
						"salary": 2.8,
						"love":   false,
						"dob":    time.Date(1990, 12, 28, 0, 0, 0, 0, time.Local).Format(model.TimeRFC3339),
					}),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Moirai{
				Storage:   &storage.SQLStorage{DB: tt.fields.DB},
				Validator: model.NewValidator(),
			}
			_, err := m.UpdateRecord(context.Background(), tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			findName := new(storage.StringCell)
			tt.fields.DB.Model(findName).Find(findName, "value = ? and record_id = ?", "dima", tt.args.record.Id)
			assert.Equal(t, findName.Value, "dima")
			assert.Equal(t, findName.Key, "name")

			findAge := new(storage.IntCell)
			tt.fields.DB.Model(findAge).Find(findAge, "value = ? and record_id = ?", 31, tt.args.record.Id)
			assert.Equal(t, findAge.Value, int64(31))
			assert.Equal(t, findAge.Key, "age")

			findSalary := new(storage.NumberCell)
			tt.fields.DB.Model(findAge).Find(findSalary, "value = ? and record_id = ?", 2.8, tt.args.record.Id)
			assert.Equal(t, findSalary.Value, 2.8)
			assert.Equal(t, findSalary.Key, "salary")

			findDOB := new(storage.DateCell)
			tt.fields.DB.Model(findAge).Find(findDOB, "record_id = ?",
				tt.args.record.Id)
			assert.Equal(t, findDOB.Value.Format(model.TimeRFC3339), time.Date(1990, 12, 28, 0, 0, 0, 0, time.Local).Format(model.TimeRFC3339))
			assert.Equal(t, findDOB.Key, "dob")

			findLove := new(storage.BoolCell)
			tt.fields.DB.Model(findAge).Find(findLove, "value = ? and record_id = ?", 0, tt.args.record.Id)
			assert.False(t, findLove.Value)
			assert.Equal(t, findLove.Key, "love")
		})
	}
}
