package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestProperties_Scan(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name      string
		c         *Properties
		args      args
		wantErr   bool
		assertKey func(p *Properties)
	}{
		{
			name: "#1: Load string data",
			c:    new(Properties),
			args: args{
				v: []byte(`{"name":{"type":"string","description":"drinking beer with Dima"}}`),
			},
			wantErr: false,
			assertKey: func(p *Properties) {
				assert.Equal(t, p.Columns["name"], &StringType{Type: StringTp, Description: "drinking beer with Dima"})
			},
		},
		{
			name: "#2: Load string and integer data",
			c:    new(Properties),
			args: args{
				v: []byte(`{"name":{"type":"string","description":"drinking beer with Dima"},"age":{"type":"integer","description":"playing game"}}`),
			},
			wantErr: false,
			assertKey: func(p *Properties) {
				assert.Equal(t, p.Columns["name"], &StringType{Type: StringTp, Description: "drinking beer with Dima"})
				assert.Equal(t, p.Columns["age"], &IntegerType{Type: IntegerTp, Description: "playing game"})
			},
		},
		{
			name: "#3: Load Float data",
			c:    new(Properties),
			args: args{
				v: []byte(`{"bet":{"type":"number","description":"betting game"}}`),
			},
			wantErr: false,
			assertKey: func(p *Properties) {
				assert.Equal(t, p.Columns["bet"], &FloatType{Type: FloatTp, Description: "betting game"})
			},
		},
		{
			name: "#1: Load Datetime data",
			c:    new(Properties),
			args: args{
				v: []byte(`{"name":{"type":"string","description":"drinking beer with Dima","format":"date-time"}}`),
			},
			wantErr: false,
			assertKey: func(p *Properties) {
				assert.Equal(t, p.Columns["name"], &DateTimeType{Type: StringTp, Format: DateTp, Description: "drinking beer with Dima"})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Scan(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.assertKey(tt.c)
		})
	}
}

func TestSchema_JSONSchema(t *testing.T) {
	type fields struct {
		ID         uuid.UUID
		Name       string
		Properties Properties
		Required   Required
		ProjectID  uuid.UUID
		Version    string
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "#1: Generate JSON schema",
			fields: fields{
				ID:   uuid.UUID{},
				Name: "Simple Schema",
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
			want: `{"type":"object","properties":{"name":{"type":"string","description":"ok","minLength":10,"maxLength":100,"default":"paul"}},"required":["name"]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Schema{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				Properties: tt.fields.Properties,
				Required:   tt.fields.Required,
				ProjectID:  tt.fields.ProjectID,
				Version:    tt.fields.Version,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			got := s.JSONSchema()
			assert.Equal(t, got, tt.want)
		})
	}
}
