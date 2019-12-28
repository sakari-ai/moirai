package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProperties_Scan(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name      string
		c         Properties
		args      args
		wantErr   bool
		assertKey func(p Properties)
	}{
		{
			name: "#1: Load string data",
			c:    make(Properties),
			args: args{
				v: []byte(`{"name":{"type":"string","description":"drinking beer with Dima"}}`),
			},
			wantErr: false,
			assertKey: func(p Properties) {
				assert.Equal(t, p["name"], &StringType{Type: StringTp, Description: "drinking beer with Dima"})
			},
		},
		{
			name: "#2: Load string and integer data",
			c:    make(Properties),
			args: args{
				v: []byte(`{"name":{"type":"string","description":"drinking beer with Dima"},"age":{"type":"integer","description":"playing game"}}`),
			},
			wantErr: false,
			assertKey: func(p Properties) {
				assert.Equal(t, p["name"], &StringType{Type: StringTp, Description: "drinking beer with Dima"})
				assert.Equal(t, p["age"], &IntegerType{Type: IntegerTp, Description: "playing game"})
			},
		},
		{
			name: "#3: Load Float data",
			c:    make(Properties),
			args: args{
				v: []byte(`{"bet":{"type":"float","description":"betting game"}}`),
			},
			wantErr: false,
			assertKey: func(p Properties) {
				assert.Equal(t, p["bet"], &FloatType{Type: FloatTp, Description: "betting game"})
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
