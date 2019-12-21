package config

import (
	"os"
	"reflect"
	"testing"

	"github.com/sakari-ai/moirai/config/loader"
	"github.com/sakari-ai/moirai/config/loader/consul"
	"github.com/sakari-ai/moirai/config/loader/file"
	"github.com/sakari-ai/moirai/config/namespace"
	"github.com/sakari-ai/moirai/config/storage"
)

func TestLoad(t *testing.T) {
	type args struct {
		storage storage.Storage
	}
	tests := []struct {
		name string
		args args
		want func(loader.Loader) bool
	}{
		{
			name: "Load local file",
			args: args{storage: "/local/file"},
			want: func(loader loader.Loader) bool {
				f, ok := loader.(*file.File)
				return ok && f.Path == "/local/file"
			}},
		{
			name: "Load consul host",
			args: args{storage: "http://consul.io/innte"},
			want: func(loader loader.Loader) bool {
				c, ok := loader.(*consul.Consul)
				return ok && c.Host == "http://consul.io/moirai"
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loader.New(tt.args.storage); !tt.want(got) {
				t.Errorf("Load: %s must get proper storage loader", tt.name)
			}
		})
	}
}

func TestGetNamespaceRegistry(t *testing.T) {
	type args struct {
		env     string
		service string
	}
	tests := []struct {
		name string
		args args
		want namespace.Namespace
	}{
		{
			name: "Create proper namesapce",
			args: args{
				env:     "stag",
				service: "kairosdb",
			},
			want: namespace.FromString("github.com/sakari-ai/moirai/stag/services/kairosdb"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("ENV", tt.args.env)
			if got := namespace.FromMode(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
