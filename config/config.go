package config

import (
	"os"
	"time"

	"github.com/sakari-ai/moirai/config/env"
	"github.com/sakari-ai/moirai/log"
	"github.com/sakari-ai/moirai/log/field"

	"github.com/sakari-ai/moirai/config/loader"
	"github.com/sakari-ai/moirai/config/namespace"
)

type Config struct {
	HTTP         *HTTP        `json:"http"`
	GRPC         *GRPC        `json:"grpc"`
	Database     *Database    `json:"database"`
	Redis        *Redis       `json:"redis"`
	Cache        *BigCache    `json:"cache"`
	BloomFilter  *BloomFilter `json:"bloomFilter"`
	GoogleOAuth2 *OAuth2      `json:"googleOauth2"`
	JWT          *JWT         `json:"jwt"`
}

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Migrate  string `json:"migrate"`
	Debug    bool   `json:"debug"`
}

type BigCache struct {
	Shards             int           `json:"shards"`
	LifeWindowMinutes  int           `json:"lifeWindowMinutes"`
	HardMaxCacheSizeMb int           `json:"hardMaxCacheSizeMb"`
	MaxEntrySize       int           `json:"maxEntrySize"`
	RefreshInSeconds   time.Duration `json:"refreshInSeconds"`
}

type Redis struct {
	Address  []string `json:"address"`
	DB       int      `json:"db"`
	Password string   `json:"password"`
}

type HTTP struct {
	Address string `json:"address"`
}

type GRPC struct {
	Address string `json:"address"`
}

type Websocket struct {
	Address string `json:"address"` // "ws://localhost:8087"
}

type BloomFilter struct {
	M uint `json:"m"`
	K uint `json:"k"`
}

type OAuth2 struct {
	ClientID     string        `json:"clientId"`
	ClientSecret string        `json:"clientSecret"`
	CallbackURL  string        `json:"callbackURL"`
	RedirectURLs []RedirectURL `json:"redirectURLs"`
}

type JWT struct {
	Secret          string        `json:"secret"`
	ExpiryInSeconds time.Duration `json:"expiryInSeconds"`
}

type RedirectURL struct {
	Default     bool   `json:"default"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

func LoadWithPlaceholder(loader loader.Loader, ph interface{}) {
	cfgNamespace := os.Getenv(env.KeyConfigNamespace)
	ns := namespace.FromString(cfgNamespace)
	err := loader.Load(ns, ph)
	if err != nil {
		log.Fatal("can not load config config",
			field.String("namespace", ns.String()),
			field.Error(err),
		)
	}
}

func LoadWithPlaceholderFromNamespace(loader loader.Loader, namespace namespace.Namespace, ph interface{}) {
	err := loader.Load(namespace, ph)
	if err != nil {
		log.Fatal("can not load config config",
			field.String("namespace", namespace.String()),
			field.Error(err),
		)
	}
}
