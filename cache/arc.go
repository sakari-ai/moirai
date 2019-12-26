package cache

import (
	"context"
	"encoding/json"
	"time"

	lru "github.com/hashicorp/golang-lru"
)

const defaultSize = 1000

type cacheRequest struct {
	Action string
	Key    interface{}
}

type DBCache struct {
	*lru.ARCCache
	context.Context
	RefreshInSeconds time.Duration
}

func New(ctx context.Context, refreshInSeconds time.Duration) *DBCache {
	a, err := lru.NewARC(defaultSize)
	if err != nil {
		panic(err)
	}
	sc := &DBCache{}
	sc.ARCCache = a
	sc.Context = ctx
	sc.RefreshInSeconds = refreshInSeconds
	return sc
}

func (s *DBCache) Get(key interface{}) (value interface{}, ok bool) {
	return s.ARCCache.Get(key)
}

func (s *DBCache) Add(key interface{}, value interface{}) {
	s.ARCCache.Add(key, value)
}

func (s *DBCache) Remove(key interface{}) {
	s.ARCCache.Remove(key)
}

func request(action string, value interface{}) []byte {
	data := cacheRequest{Action: action, Key: value}

	r, _ := json.Marshal(data)
	return r
}
