package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	contents  []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: map[string]cacheEntry{},
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		contents:  val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	fmt.Println("Checking cache for: ", key)
	for key, val := range c.cache {
		fmt.Printf("> %v (%v)\n", key, val.createdAt)
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	if v, exists := c.cache[key]; exists {
		fmt.Println("Loading existing data...")
		return v.contents, true
	}
	fmt.Println("No data found in cache...")
	return []byte{}, false
}

func (c *Cache) Count() int {
	return len(c.cache)
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
	fmt.Println("Running check for old cache data...")
	for key, val := range c.cache {
		if val.createdAt.Before(now.Add(-last)) {
			fmt.Println("Removing ", key)
			delete(c.cache, key)
		}
	}
}
