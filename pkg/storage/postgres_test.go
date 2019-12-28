package storage

import (
	"github.com/sakari-ai/moirai/database"
	"github.com/sakari-ai/moirai/database/fake"
	"github.com/sakari-ai/moirai/pkg/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
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
			p := PostgresStorage{
				DB: tt.fields.DB,
			}
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
