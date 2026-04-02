package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	mu    sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{}
	cache.cache = make(map[string]CacheEntry)
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := CacheEntry{}
	entry.val = val
	entry.createdAt = time.Now()
	c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, present := c.cache[key]
	if !present {
		return nil, false

	}
	return entry.val, true

}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for t := range ticker.C {
		c.mu.Lock()
		fmt.Println("Clearing old items from cache")
		for key, val := range c.cache {
			if t.UTC().Sub(val.createdAt) > interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()

	}
}
