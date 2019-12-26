package cache

type Cache interface {
	Get(key interface{}) (value interface{}, ok bool)
	Add(key interface{}, value interface{})
	Remove(key interface{})
}
