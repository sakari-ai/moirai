package model

import (
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/sakari-ai/moirai/core/util"
	uuid "github.com/satori/go.uuid"
	"github.com/xeipuuv/gojsonschema"
	"sync"
	"testing"
	"time"
)

func TestJsonSchemaValidator_Validate(t *testing.T) {
	t.Parallel()
	type fields struct {
		schemaLoader map[string]*gojsonschema.Schema
		RWMutex      *sync.RWMutex
	}
	type args struct {
		schema Schema
		record *structpb.Struct
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "#1: TRUE - Validate Simple string schema",
			fields: fields{
				schemaLoader: map[string]*gojsonschema.Schema{},
				RWMutex:      new(sync.RWMutex),
			},
			args: args{
				schema: Schema{
					ID:      uuid.UUID{},
					Version: "12341234123413",
					Name:    "Simple Schema",
					Properties: Properties{Columns: map[string]PropertyType{
						"name": &StringType{
							Type:        "string",
							Description: "ok",
							Default:     "paul",
							MinLength:   2,
							MaxLength:   100,
						},
					}},
					Required: []string{"name"},
				},
				record: util.StructProto(map[string]interface{}{
					"name": "Paul Aan",
				}),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "#2: FALSE - Validate Simple string schema",
			fields: fields{
				schemaLoader: map[string]*gojsonschema.Schema{},
				RWMutex:      new(sync.RWMutex),
			},
			args: args{
				schema: Schema{
					ID:      uuid.UUID{},
					Version: "12341234123413",
					Name:    "Simple Schema",
					Properties: Properties{Columns: map[string]PropertyType{
						"name": &StringType{
							Type:        "string",
							Description: "ok",
							Default:     "paul",
							MinLength:   10,
							MaxLength:   100,
						},
					}},
					Required: []string{"name"},
				},
				record: util.StructProto(map[string]interface{}{
					"name": "Paul",
				}),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "#3: TRUE - Validate Simple Integer schema",
			fields: fields{
				schemaLoader: map[string]*gojsonschema.Schema{},
				RWMutex:      new(sync.RWMutex),
			},
			args: args{
				schema: Schema{
					ID:      uuid.UUID{},
					Version: "12341234123413",
					Name:    "Simple Schema",
					Properties: Properties{Columns: map[string]PropertyType{
						"age": &IntegerType{
							Type:        IntegerTp,
							Description: "integer",
							Default:     10,
							Minimum:     10,
							Maximum:     100,
						},
					}},
					Required: []string{"age"},
				},
				record: util.StructProto(map[string]interface{}{
					"age": 10,
				}),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "#4: FALSE - Validate Simple Integer schema",
			fields: fields{
				schemaLoader: map[string]*gojsonschema.Schema{},
				RWMutex:      new(sync.RWMutex),
			},
			args: args{
				schema: Schema{
					ID:      uuid.UUID{},
					Version: "12341234123413",
					Name:    "Simple Schema",
					Properties: Properties{Columns: map[string]PropertyType{
						"age": &IntegerType{
							Type:        IntegerTp,
							Description: "integer",
							Default:     10,
							Minimum:     10,
							Maximum:     100,
						},
					}},
					Required: []string{"age"},
				},
				record: util.StructProto(map[string]interface{}{
					"age": 5,
				}),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "#5: TRUE - Validate Simple DateTime schema",
			fields: fields{
				schemaLoader: map[string]*gojsonschema.Schema{},
				RWMutex:      new(sync.RWMutex),
			},
			args: args{
				schema: Schema{
					ID:      uuid.UUID{},
					Version: "12341234123413",
					Name:    "Simple Schema",
					Properties: Properties{Columns: map[string]PropertyType{
						"dob": &DateTimeType{
							Type:        StringTp,
							Description: "date",
							Default:     time.Date(2019, 12, 28, 0, 0, 0, 0, time.Local).Format(TimeRFC3339),
							Format:      "date-time",
						},
					}},
					Required: []string{"dob"},
				},
				record: util.StructProto(map[string]interface{}{
					"dob": "2019-12-29T08:30:06.283185Z",
				}),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "#6: FALSE - Validate Simple DateTime schema - wrong datetime format",
			fields: fields{
				schemaLoader: map[string]*gojsonschema.Schema{},
				RWMutex:      new(sync.RWMutex),
			},
			args: args{
				schema: Schema{
					ID:      uuid.UUID{},
					Version: "12341234123413",
					Name:    "Simple Schema",
					Properties: Properties{Columns: map[string]PropertyType{
						"dob": &DateTimeType{
							Type:        StringTp,
							Description: "date",
							Default:     time.Date(2019, 12, 28, 0, 0, 0, 0, time.Local).Format(TimeRFC3339),
							Format:      "date-time",
						},
					}},
					Required: []string{"dob"},
				},
				record: util.StructProto(map[string]interface{}{
					"dob": "201912-27T08:30:06.283185Z",
				}),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "#7: TRUE - Validate Simple Boolean schema",
			fields: fields{
				schemaLoader: map[string]*gojsonschema.Schema{},
				RWMutex:      new(sync.RWMutex),
			},
			args: args{
				schema: Schema{
					ID:      uuid.UUID{},
					Version: "12341234123413",
					Name:    "Simple Schema",
					Properties: Properties{Columns: map[string]PropertyType{
						"dob": &BooleanType{
							Type:        BooleanTp,
							Description: "bool",
						},
					}},
					Required: []string{"dob"},
				},
				record: util.StructProto(map[string]interface{}{
					"dob": true,
				}),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "#8: False - Validate Simple Boolean schema",
			fields: fields{
				schemaLoader: map[string]*gojsonschema.Schema{},
				RWMutex:      new(sync.RWMutex),
			},
			args: args{
				schema: Schema{
					ID:      uuid.UUID{},
					Version: "12341234123413",
					Name:    "Simple Schema",
					Properties: Properties{Columns: map[string]PropertyType{
						"dob": &BooleanType{
							Type:        BooleanTp,
							Description: "bool",
						},
					}},
					Required: []string{"dob"},
				},
				record: util.StructProto(map[string]interface{}{
					"dob": "x",
				}),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "#9: TRUE - Validate Simple Float schema",
			fields: fields{
				schemaLoader: map[string]*gojsonschema.Schema{},
				RWMutex:      new(sync.RWMutex),
			},
			args: args{
				schema: Schema{
					ID:      uuid.UUID{},
					Version: "12341234123413",
					Name:    "Simple Schema",
					Properties: Properties{Columns: map[string]PropertyType{
						"age": &FloatType{
							Type:        FloatTp,
							Description: "integer",
							Default:     10.0,
							Minimum:     10.0,
							Maximum:     100.0,
						},
					}},
					Required: []string{"age"},
				},
				record: util.StructProto(map[string]interface{}{
					"age": 11.0,
				}),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "#10: FALSE - Validate Simple FloatType schema -  Wrong format",
			fields: fields{
				schemaLoader: map[string]*gojsonschema.Schema{},
				RWMutex:      new(sync.RWMutex),
			},
			args: args{
				schema: Schema{
					ID:      uuid.UUID{},
					Version: "12341234123413",
					Name:    "Simple Schema",
					Properties: Properties{Columns: map[string]PropertyType{
						"age": &FloatType{
							Type:        FloatTp,
							Description: "integer",
							Default:     10.0,
							Minimum:     10.0,
							Maximum:     100.0,
						},
					}},
					Required: []string{"age"},
				},
				record: util.StructProto(map[string]interface{}{
					"age": "9",
				}),
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator()
			got, err := v.Validate(tt.args.schema, tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Validate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
