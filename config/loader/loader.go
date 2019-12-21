package loader

import (
	"net/url"

	"github.com/sakari-ai/moirai/config/loader/consul"
	"github.com/sakari-ai/moirai/config/loader/file"
	"github.com/sakari-ai/moirai/config/namespace"
	"github.com/sakari-ai/moirai/config/storage"
)

type Loader interface {
	Load(namespace namespace.Namespace, value interface{}) error
}

func New(s storage.Storage) Loader {
	storage := s.String()
	url, err := url.ParseRequestURI(storage)
	if err != nil || (url.Host == "" && url.Scheme == "") {
		return file.New(storage)
	}
	return consul.New(storage)
}
