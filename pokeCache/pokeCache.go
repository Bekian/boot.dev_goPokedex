package pokecache

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.Mutex
	cache    map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) (cache *Cache) {
	cacheMap := make(map[string]cacheEntry)
	newInterval := time.Duration(interval) * time.Second
	cache = &Cache{cache: cacheMap, interval: newInterval}
	go cache.ReapLoop()
	return
}

func (c *Cache) Add(key string, val []byte) {
	createdAt := time.Now()
	c.cache[key] = cacheEntry{createdAt: createdAt, val: val}
}

func (c *Cache) Get(key string) (val []byte, exists bool) {
	var entry cacheEntry
	entry, exists = c.cache[key]
	val = entry.val
	return
}

func (c *Cache) ReapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C { // Waits for the next tick
		now := time.Now()

		// Lock the cache to safely access shared data
		c.mu.Lock()
		for key, entry := range c.cache {
			if now.Sub(entry.createdAt) > c.interval {
				delete(c.cache, key)
				fmt.Println(now.Sub(entry.createdAt), c.interval)
				os.Exit(0)
			}
		}
		c.mu.Unlock()
	}
}
