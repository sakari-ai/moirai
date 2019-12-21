package file

import (
	"os"
	"path"
	"testing"

	"github.com/sakari-ai/moirai/config/namespace"

	"github.com/stretchr/testify/assert"

	_ "github.com/spf13/viper/remote"
)

func TestFile_Success(t *testing.T) {
	type config struct {
		Status int `json:"status"`
	}
	type args struct {
		namespace namespace.Namespace
	}

	var (
		cfg    = &config{}
		pwd, _ = os.Getwd()
		tests  = []struct {
			name    string
			path    string
			args    args
			wantErr bool
		}{
			{
				name:    "Load from existing file",
				path:    path.Join(pwd, "testdata/mock.yaml"),
				args:    args{namespace: "moirai"},
				wantErr: false,
			},
			{
				name:    "Load from nonexistent file",
				path:    path.Join(pwd, "testdata/unknown.yaml"),
				args:    args{namespace: "moirai"},
				wantErr: true,
			},
		}
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.path)
			if err := v.Load(tt.args.namespace, cfg); (err != nil) != tt.wantErr {
				t.Errorf("local.Load() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Equal(t, 1, cfg.Status, "Must read data from mock.yaml")
			}
		})
	}
}
