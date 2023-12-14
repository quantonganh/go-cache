package cache

import (
	"sync"
	"time"
)

type Cache struct {
	m sync.Map
}

func New() *Cache {
	return &Cache{}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.m.Store(key, value)
	time.AfterFunc(ttl, func() {
		c.m.Delete(key)
	})
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.m.Load(key)
}
