package consul

import (
	"testing"

	"github.com/sakari-ai/moirai/config/namespace"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

const consulValidResponse = `[{"LockIndex":0,"Key":"trackclix/exp","Flags":0,"Value":"eyJzdGF0dXMiOiAxfQ==","CreateIndex":1716,"ModifyIndex":2140}]`

func TestConsul_Success(t *testing.T) {
	type config struct {
		Status int `json:"status"`
	}
	type args struct {
		namespace namespace.Namespace
		host      string
	}

	// mock the expected outgoing request for new config
	defer gock.Off()

	// mock consul server
	gock.New("http://consul.io").
		Get("/v1/kv/trackclix/exp").
		Reply(200).
		Type("json").
		BodyString(consulValidResponse)

	var (
		cfg   = &config{}
		tests = []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: "Read consul having Innete",
				args: args{
					namespace: "trackclix/exp",
					host:      "http://consul.io",
				},
				wantErr: false,
			},
		}
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.args.host)
			if err := v.Load(tt.args.namespace, cfg); (err != nil) != tt.wantErr {
				t.Errorf("local.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, 1, cfg.Status, "Must reflect 1")
		})
	}
}

func TestConsul_Error(t *testing.T) {
	type config struct {
		Status int `json:"status"`
	}
	type args struct {
		namespace namespace.Namespace
		host      string
	}

	// mock the expected outgoing request for new config
	defer gock.Off()

	// mock consul server
	gock.New("http://consul.io").
		Get("/v1/kv/trackclix/exp").
		Reply(200).
		Type("json").
		BodyString(``)

	var (
		cfg   = &config{}
		tests = []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: "Read consul having Innete",
				args: args{
					namespace: "trackclix/exp",
					host:      "http://consul.io",
				},
				wantErr: true,
			},
		}
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.args.host)
			if err := v.Load(tt.args.namespace, cfg); (err != nil) != tt.wantErr {
				t.Errorf("local.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, 0, cfg.Status, "Must reflect 1")
		})
	}
}
