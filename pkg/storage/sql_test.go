package storage

import (
	"errors"
	"github.com/sakari-ai/moirai/database"
	"github.com/sakari-ai/moirai/database/fake"
	"github.com/sakari-ai/moirai/pkg/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestPostgresStorage_GetSchema(t *testing.T) {
	type fields struct {
		DB database.DBEngine
	}
	type args struct {
		item *model.Schema
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Schema
		wantErr bool
	}{
		{
			name: "#1: Read and Write",
			fields: fields{DB: func() database.DBEngine {
				eng := fake.PrepareForTesting(t, fake.WithModels(&model.Schema{}))

				return eng
			}()},
			args: args{
				item: &model.Schema{
					ID:   uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57c"),
					Name: "Writing",
					Properties: model.Properties{
						Columns: map[string]model.PropertyType{
							"name": &model.StringType{
								Type:        "string",
								Description: "write me",
								MinLength:   1,
								MaxLength:   10,
								Default:     "ok",
							},
						},
					},
					Required:  []string{"matchId"},
					ProjectID: uuid.NewV4(),
					Version:   "xxxx",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.fields.DB.Close()
			p := NewStorage()
			p.DB = tt.fields.DB
			p.WriteSchema(tt.args.item)
			got, err := p.GetSchema(tt.args.item.ID)
			assert.ObjectsAreEqualValues(got.Properties, model.Properties{
				Columns: map[string]model.PropertyType{
					"name": &model.StringType{
						Type:        "string",
						Description: "write me",
						MinLength:   1,
						MaxLength:   10,
						Default:     "ok",
					},
				},
			})
			assert.Equal(t, got.Required, model.Required{"matchId"}, "Must contain previous mock")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPostgresStorage_WriteRecord(t *testing.T) {
	type fields struct {
		DB database.DBEngine
	}
	type args struct {
		schema model.Schema
		record *model.Record
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Record
		wantErr bool
	}{
		{
			name: "#1: Write simple schema",
			fields: fields{DB: func() database.DBEngine {
				mockDB := fake.PrepareForTesting(t, fake.WithModels(&model.Schema{}, &model.Record{}, &IntCell{}, &NumberCell{}, &StringCell{}, &DateCell{}, &BoolCell{}))

				return mockDB
			}()},
			args: args{
				schema: model.Schema{
					Properties: model.Properties{
						Columns: map[string]model.PropertyType{
							"name":   &model.StringType{Type: model.StringTp},
							"age":    &model.IntegerType{Type: model.IntegerTp},
							"salary": &model.FloatType{Type: model.FloatTp},
							"love":   &model.BooleanType{Type: model.BooleanTp},
							"dob":    &model.DateTimeType{Type: model.StringTp, Format: model.DateTp},
						}},
				},
				record: &model.Record{
					ID:        uuid.UUID{},
					ProjectID: uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57c"),
					SchemaID:  uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57b"),
					Fields: model.Fields{
						Columns: map[string]interface{}{
							"name":   "paul",
							"age":    35,
							"salary": 36.6,
							"love":   true,
							"dob":    time.Date(1986, 12, 28, 0, 0, 0, 0, time.Local).Format(model.TimeRFC3339),
						},
					},
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
			},
			want: &model.Record{
				ProjectID: uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57c"),
				SchemaID:  uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57b"),
				Fields: model.Fields{
					Columns: map[string]interface{}{
						"name":   "paul",
						"age":    35,
						"salary": 36.6,
						"love":   true,
						"dob":    time.Date(1986, 12, 28, 0, 0, 0, 0, time.Local).Format(model.TimeRFC3339),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SQLStorage{
				DB: tt.fields.DB,
			}
			defer tt.fields.DB.Close()
			got, err := p.WriteRecord(tt.args.schema, tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want.ProjectID, got.ProjectID)
			assert.Equal(t, tt.want.SchemaID, got.SchemaID)
			assert.Equal(t, tt.want.Fields, got.Fields)

			findAge := new(IntCell)
			tt.fields.DB.Model(findAge).Find(findAge, "value = ? and record_id = ?", 35, tt.args.record.ID)
			assert.Equal(t, findAge.Value, int64(35))
			assert.Equal(t, findAge.Key, "age")

			findName := new(StringCell)
			tt.fields.DB.Model(findName).Find(findName, "value = ? and record_id = ?", "paul", tt.args.record.ID)
			assert.Equal(t, findName.Value, "paul")
			assert.Equal(t, findName.Key, "name")

			findSalary := new(NumberCell)
			tt.fields.DB.Model(findAge).Find(findSalary, "value = ? and record_id = ?", 36.6, tt.args.record.ID)
			assert.Equal(t, findSalary.Value, 36.6)
			assert.Equal(t, findSalary.Key, "salary")

			findDOB := new(DateCell)
			tt.fields.DB.Model(findAge).Find(findDOB, "record_id = ?",
				tt.args.record.ID)
			assert.Equal(t, findDOB.Value.Format(model.TimeRFC3339), time.Date(1986, 12, 28, 0, 0, 0, 0, time.Local).Format(model.TimeRFC3339))
			assert.Equal(t, findDOB.Key, "dob")

			findLove := new(BoolCell)
			tt.fields.DB.Model(findAge).Find(findLove, "value = ? and record_id = ?", true, tt.args.record.ID)
			assert.True(t, findLove.Value)
			assert.Equal(t, findLove.Key, "love")
		})
	}
}

type mockDbEngine struct {
	mock.Mock
	database.DBEngine
}

func (m *mockDbEngine) Create(value interface{}) database.DBEngine {
	arg := m.Called(value)
	return arg.Get(0).(database.DBEngine)
}

func (m *mockDbEngine) Error() error {
	arg := m.Called()
	return arg.Error(0)
}

func (m *mockDbEngine) Rollback() database.DBEngine {
	arg := m.Called()
	return arg.Get(0).(database.DBEngine)
}

func (m *mockDbEngine) Begin() database.DBEngine {
	arg := m.Called()
	return arg.Get(0).(database.DBEngine)
}

func (m *mockDbEngine) Commit() database.DBEngine {
	arg := m.Called()
	return arg.Get(0).(database.DBEngine)
}

func TestPostgresStorage_WriteRecord_Error(t *testing.T) {
	type fields struct {
		DB database.DBEngine
	}
	type args struct {
		schema model.Schema
		record *model.Record
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Record
		wantErr bool
	}{
		{
			name: "#1: Write simple schema",
			fields: fields{DB: func() database.DBEngine {
				mockDB := new(mockDbEngine)
				mockDB.On("Create", mock.Anything).Return(mockDB)
				mockDB.On("Rollback").Return(mockDB)
				mockDB.On("Begin").Return(mockDB)
				mockDB.On("Commit").Return(mockDB)
				mockDB.On("Error").Return(errors.New("db error"))
				return mockDB
			}()},
			args: args{
				schema: model.Schema{
					Properties: model.Properties{
						Columns: map[string]model.PropertyType{
							"name":   &model.StringType{Type: model.StringTp},
							"age":    &model.IntegerType{Type: model.IntegerTp},
							"salary": &model.FloatType{Type: model.FloatTp},
							"love":   &model.BooleanType{Type: model.BooleanTp},
							"dob":    &model.DateTimeType{Type: model.StringTp, Format: model.DateTp},
						}},
				},
				record: &model.Record{
					ID:        uuid.UUID{},
					ProjectID: uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57c"),
					SchemaID:  uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57b"),
					Fields: model.Fields{
						Columns: map[string]interface{}{
							"name":   "paul",
							"age":    35,
							"salary": 36.6,
							"love":   true,
							"dob":    time.Date(1986, 12, 28, 0, 0, 0, 0, time.Local).Format(model.TimeRFC3339),
						},
					},
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
			},
			want: &model.Record{
				ProjectID: uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57c"),
				SchemaID:  uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57b"),
				Fields: model.Fields{
					Columns: map[string]interface{}{
						"name":   "paul",
						"age":    35,
						"salary": 36.6,
						"love":   true,
						"dob":    time.Date(1986, 12, 28, 0, 0, 0, 0, time.Local).Format(model.TimeRFC3339),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SQLStorage{
				DB: tt.fields.DB,
			}
			_, err := p.WriteRecord(tt.args.schema, tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSQLStorage_UpdateRecord(t *testing.T) {
	sch := &model.Schema{
		Name: "Simple One",
		Properties: model.Properties{
			Columns: map[string]model.PropertyType{
				"name": &model.StringType{
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
		DB database.DBEngine
	}
	type args struct {
		schema model.Schema
		record *model.Record
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
					mockDB := fake.PrepareForTesting(t, fake.WithModels(&model.Schema{}, &model.Record{}, &IntCell{}, &NumberCell{}, &StringCell{}, &DateCell{}, &BoolCell{}))

					mockDB.Create(sch)
					mockDB.Create(rcd)
					mockDB.Create(&StringCell{
						RecordID: rcd.ID,
						Key:      "name",
						Value:    "paul",
					})
					mockDB.Create(&IntCell{
						RecordID: rcd.ID,
						Key:      "age",
						Value:    35,
					})
					mockDB.Create(&NumberCell{
						RecordID: rcd.ID,
						Key:      "salary",
						Value:    3.5,
					})
					mockDB.Create(&BoolCell{
						RecordID: rcd.ID,
						Key:      "love",
						Value:    true,
					})
					mockDB.Create(&DateCell{
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
				record: &model.Record{
					ID:        rcd.ID,
					ProjectID: uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57c"),
					SchemaID:  sch.ID,
					Fields: model.Fields{
						Columns: map[string]interface{}{
							"name":   "dima",
							"age":    31,
							"salary": 2.8,
							"love":   false,
							"dob":    time.Date(1990, 12, 28, 0, 0, 0, 0, time.Local).Format(model.TimeRFC3339),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SQLStorage{
				DB: tt.fields.DB,
			}
			_, err := p.UpdateRecord(tt.args.schema, tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			findName := new(StringCell)
			tt.fields.DB.Model(findName).Find(findName, "value = ? and record_id = ?", "dima", tt.args.record.ID)
			assert.Equal(t, findName.Value, "dima")
			assert.Equal(t, findName.Key, "name")

			findAge := new(IntCell)
			tt.fields.DB.Model(findAge).Find(findAge, "value = ? and record_id = ?", 31, tt.args.record.ID)
			assert.Equal(t, findAge.Value, int64(31))
			assert.Equal(t, findAge.Key, "age")

			findSalary := new(NumberCell)
			tt.fields.DB.Model(findAge).Find(findSalary, "value = ? and record_id = ?", 2.8, tt.args.record.ID)
			assert.Equal(t, findSalary.Value, 2.8)
			assert.Equal(t, findSalary.Key, "salary")

			findDOB := new(DateCell)
			tt.fields.DB.Model(findAge).Find(findDOB, "record_id = ?",
				tt.args.record.ID)
			assert.Equal(t, findDOB.Value.Format(model.TimeRFC3339), time.Date(1990, 12, 28, 0, 0, 0, 0, time.Local).Format(model.TimeRFC3339))
			assert.Equal(t, findDOB.Key, "dob")

			findLove := new(BoolCell)
			tt.fields.DB.Model(findAge).Find(findLove, "value = ? and record_id = ?", 0, tt.args.record.ID)
			assert.False(t, findLove.Value)
			assert.Equal(t, findLove.Key, "love")
		})
	}
}
