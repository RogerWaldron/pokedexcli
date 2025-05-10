package pokecache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.RWMutex
}

type cacheEntry struct {
	createdAt	time.Time
	val         []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		 cache: make(map[string]cacheEntry),
		 mux: &sync.RWMutex{},
	}

	go cache.reapLoop(interval)

	return cache
} 

func(c *Cache) Add(key string, value []byte) error{
	if key == "" {
		return errors.New("error Adding to cache: key missing")
	}

	entry := cacheEntry{
		createdAt: time.Now().UTC(),
		val: value,
	}

	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = entry

	return nil
}

func(c *Cache) Get(key string) (value []byte, found bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	entry, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	} 

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}