package consul

import (
	"net/url"
	"time"

	"github.com/mitchellh/mapstructure"

	"github.com/sakari-ai/moirai/config/namespace"

	"github.com/spf13/viper"
	// enable load from remote
	_ "github.com/spf13/viper/remote"
)

type Consul struct {
	Host string
}

func New(host string) *Consul {
	return &Consul{
		Host: host,
	}
}

func (c *Consul) Load(namespace namespace.Namespace, value interface{}) error {
	url, err := url.Parse(c.Host)
	v := viper.New()
	err = v.AddRemoteProvider("consul", url.Host, namespace.String())
	if err != nil {
		return err
	}
	v.SetConfigType("json")

	err = v.ReadRemoteConfig()
	if err != nil {
		return err
	}
	return v.Unmarshal(value, viper.DecodeHook(mapstructure.StringToTimeHookFunc(time.RFC3339Nano)))
}
