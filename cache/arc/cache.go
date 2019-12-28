package arc

import (
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto"
)

type BigCache struct {
	Shards             int           `json:"shards"`
	LifeWindowMinutes  int           `json:"lifeWindowMinutes"`
	HardMaxCacheSizeMb int           `json:"hardMaxCacheSizeMb"`
	MaxEntrySize       int           `json:"maxEntrySize"`
	RefreshInSeconds   time.Duration `json:"refreshInSeconds"`
}

type Cache interface {
	Get(key interface{}) (value interface{}, ok bool)
	Add(key interface{}, value interface{})
	Remove(key interface{})
}

type dbCache struct {
	cache *ristretto.Cache
}

func (d *dbCache) Get(key interface{}) (value interface{}, ok bool) {
	return d.cache.Get(fmt.Sprint(key))
}

func (d *dbCache) Add(key interface{}, value interface{}) {
	d.cache.Set(fmt.Sprint(key), value, 1)
}

func (d *dbCache) Remove(key interface{}) {
	d.cache.Del(fmt.Sprint(key))
}

func NewCache(size int64) *dbCache {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1000000 * 10,
		MaxCost:     1000000,
		BufferItems: size,
	})
	if err != nil {
		panic(err)
	}
	return &dbCache{cache}
}
