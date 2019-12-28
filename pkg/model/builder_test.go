package model

import (
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/sakari-ai/moirai/core/util"
	"github.com/sakari-ai/moirai/internal/lib/ptypes"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateProperty(t *testing.T) {
	type args struct {
		p *DTOStruct
	}
	tests := []struct {
		name    string
		args    args
		want    PropertyType
		wantErr bool
	}{
		{
			name: "#1: String Type",
			args: args{p: func() *DTOStruct {
				jsonSchema := util.StructProto(map[string]interface{}{
					"type":        "string",
					"default":     "paul",
					"description": "description str",
					"minLength":   1,
					"maxLength":   10})
				dto := DTOStruct(*jsonSchema)

				return &dto
			}()},
			want: &StringType{
				Description: "description str",
				MinLength:   1,
				MaxLength:   10,
				Default:     "paul",
				Type:        "string",
			},
			wantErr: false,
		},
		{
			name: "#2: Integer Type",
			args: args{p: func() *DTOStruct {
				jsonSchema := util.StructProto(map[string]interface{}{
					"type":        "integer",
					"default":     10,
					"description": "description int",
					"minimum":     1,
					"maximum":     10})
				dto := DTOStruct(*jsonSchema)

				return &dto
			}()},
			want: &IntegerType{
				Description: "description int",
				Minimum:     1,
				Maximum:     10,
				Default:     10,
				Type:        "integer",
			},
			wantErr: false,
		},
		{
			name: "#3: Float Type",
			args: args{p: func() *DTOStruct {
				jsonSchema := util.StructProto(map[string]interface{}{
					"type":        "float",
					"default":     2.2,
					"description": "description float",
					"minimum":     1.0,
					"maximum":     10.0})
				dto := DTOStruct(*jsonSchema)

				return &dto
			}()},
			want: &FloatType{
				Description: "description float",
				Minimum:     1.0,
				Maximum:     10.0,
				Default:     2.2,
				Type:        "float",
			},
			wantErr: false,
		},
		{
			name: "#4: DateTime Type",
			args: args{p: func() *DTOStruct {
				jsonSchema := util.StructProto(map[string]interface{}{
					"type":        "date",
					"default":     time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local).Format(time.RFC3339),
					"description": "description date",
					"minimum":     time.Date(2019, 12, 28, 0, 0, 0, 0, time.Local).Format(time.RFC3339),
					"maximum":     time.Date(2019, 12, 30, 0, 0, 0, 0, time.Local).Format(time.RFC3339)})
				dto := DTOStruct(*jsonSchema)

				return &dto
			}()},
			want: &DateTimeType{
				Description: "description date",
				Minimum:     ptypes.TimestampProto(time.Date(2019, 12, 28, 0, 0, 0, 0, time.Local)),
				Maximum:     ptypes.TimestampProto(time.Date(2019, 12, 30, 0, 0, 0, 0, time.Local)),
				Default:     ptypes.TimestampProto(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)),
				Type:        "date",
			},
			wantErr: false,
		},
		{
			name: "#5: Boolean Type",
			args: args{p: func() *DTOStruct {
				jsonSchema := util.StructProto(map[string]interface{}{
					"type":        "boolean",
					"description": "boolean description",
				})
				dto := DTOStruct(*jsonSchema)

				return &dto
			}()},
			want: &BooleanType{
				Description: "boolean description",
				Type:        "boolean",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateProperty(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProperty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.ObjectsAreEqual(got, tt.want) {
				t.Errorf("CreateProperty() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSchema(t *testing.T) {
	type args struct {
		name      string
		projectId string
		columns   map[string]*structpb.Struct
	}
	tests := []struct {
		name    string
		args    args
		want    *Schema
		wantErr bool
	}{
		{
			name: "#1: Create simple Schema",
			args: args{
				name:      "Simple One",
				projectId: "3341fa1e-90b0-482a-b0ac-74a76d6af57c",
				columns: Columns{
					"name": util.StructProto(map[string]interface{}{
						"type":        "string",
						"default":     "paul",
						"description": "description str",
						"minLength":   1,
						"maxLength":   10}),
				},
			},
			want: &Schema{
				Name: "Simple One",
				Properties: Properties{
					Columns: map[string]PropertyType{
						"name": &StringType{
							Description: "description str",
							MinLength:   1,
							MaxLength:   10,
							Default:     "paul",
						},
					},
				},
				ProjectID: uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57c"),
			},
			wantErr: false,
		},
		{
			name: "#2: Create simple Schema with Int and String",
			args: args{
				name:      "Simple One",
				projectId: "3341fa1e-90b0-482a-b0ac-74a76d6af57c",
				columns: Columns{
					"name": util.StructProto(map[string]interface{}{
						"type":        "string",
						"default":     "paul",
						"description": "description str",
						"minLength":   1,
						"maxLength":   10}),
					"age": util.StructProto(map[string]interface{}{
						"type":        "integer",
						"default":     10,
						"description": "description int",
						"minimum":     18,
						"maximum":     60}),
				},
			},
			want: &Schema{
				Name: "Simple One",
				Properties: Properties{
					Columns: map[string]PropertyType{
						"name": &StringType{
							Description: "description str",
							MinLength:   1,
							MaxLength:   10,
							Default:     "paul",
						},
						"age": &IntegerType{
							Description: "description int",
							Minimum:     18,
							Maximum:     60,
							Default:     10,
						},
					},
				},
				ProjectID: uuid.FromStringOrNil("3341fa1e-90b0-482a-b0ac-74a76d6af57c"),
			},
			wantErr: false,
		},
		{
			name: "#2: Create simple Schema with Error not found data type",
			args: args{
				name:      "Error",
				projectId: "3341fa1e-90b0-482a-b0ac-74a76d6af57c",
				columns: Columns{
					"name": util.StructProto(map[string]interface{}{
						"type":        "double",
						"default":     "paul",
						"description": "description str",
						"minLength":   1,
						"maxLength":   10}),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSchema(tt.args.name, tt.args.projectId, tt.args.columns)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.Equal(t, got.ProjectID, tt.want.ProjectID)
				assert.ObjectsAreEqualValues(got.Properties, tt.want.Properties)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
