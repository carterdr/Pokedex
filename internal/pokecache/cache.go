package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheMap map[string]cacheEntry
	mutex    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{cacheMap: make(map[string]cacheEntry), mutex: &sync.Mutex{}}
	go c.reaploop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry := cacheEntry{createdAt: time.Now().UTC(), val: val}
	c.cacheMap[key] = entry
}

func (c *Cache) Get(key string) (val []byte, found bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, exists := c.cacheMap[key]
	return entry.val, exists
}

func (c *Cache) reaploop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.cacheMap {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cacheMap, k)
		}
	}
}
