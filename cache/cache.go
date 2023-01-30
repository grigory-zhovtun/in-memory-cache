package cache

import (
	"sync"
)

type InMemoryCache struct {
	mu    sync.Mutex
	cache map[string]interface{}
}

func New() *InMemoryCache {
	return &InMemoryCache{
		cache: make(map[string]interface{}),
	}
}

func (c *InMemoryCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c *InMemoryCache) Get(key string) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cache[key]
}

func (c *InMemoryCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.cache, key)
}