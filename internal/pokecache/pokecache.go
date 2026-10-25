package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       *sync.Mutex
	v        map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {

	cache := &Cache{
		mu:       &sync.Mutex{},
		v:        map[string]cacheEntry{},
		interval: interval,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newEntry := cacheEntry{
		val:       value,
		createdAt: time.Now(),
	}
	c.v[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.v[key]
	return value.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for range ticker.C {
		c.mu.Lock()
		for key := range c.v {
			elapsed := time.Since(c.v[key].createdAt)
			if elapsed > c.interval {
				delete(c.v, key)
			}
		}
		c.mu.Unlock()

	}

}
