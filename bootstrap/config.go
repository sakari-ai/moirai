package bootstrap

import (
	"os"
	"sync"

	"github.com/sakari-ai/moirai/config/env"
	"github.com/sakari-ai/moirai/config/loader"
	"github.com/sakari-ai/moirai/config/storage"
)

var (
	once      sync.Once
	cfgLoader loader.Loader
)

func initialize() {
	once.Do(func() {
		s := storage.FromString(os.Getenv(env.KeyConfigStorage))

		cfgLoader = loader.New(s)
	})
}

func init() {
	initialize()
}
