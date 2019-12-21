package file

import (
	"github.com/sakari-ai/moirai/config/namespace"
	"github.com/spf13/viper"
)

type File struct {
	Path string
}

func New(path string) *File {
	return &File{
		Path: path,
	}
}

func (f *File) Load(namespace namespace.Namespace, value interface{}) error {
	viper.SetConfigFile(f.Path)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.UnmarshalKey(namespace.String(), value)
}
