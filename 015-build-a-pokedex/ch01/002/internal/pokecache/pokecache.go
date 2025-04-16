package pokecache

import (
	"sync"
	"time"
)

// Cache -
type Cache struct {
	values map[string]cacheEntry
	mu     *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Add -
func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.values[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

// Get -
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, exists := c.values[key]
	return value.val, exists
}

// reapLoop -
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.values {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.values, k)
		}
	}
}

// NewCache -
func NewCache(interval time.Duration) Cache {
	c := Cache{
		values: make(map[string]cacheEntry),
		mu:     &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}
