package caching

import (
	"sync"
)

type CacheManager[T any] struct {
	Items map[string]CacheEntry[T]
	Lock  sync.Mutex
}

// Creates a new instance of cache manager
func CreateNewCacheManager[T any]() *CacheManager[T] {
	return &CacheManager[T]{}
}

// Sets or updates the wanted item
func (c *CacheManager[T]) SetOrUpdateItem(key string, item T) {

	c.Lock.Lock()
	defer c.Lock.Unlock()

	entry := CreateNewCacheEntry[T](item)
	c.Items[key] = entry
}

// Checks if a key exists
func (c *CacheManager[T]) keyExists(key string) bool {
	for k := range c.Items {
		if k == key {
			return true
		}
	}
	return false
}

// Returns a item by its key
func (c *CacheManager[T]) GetItem(key string) (T, bool) {

	c.Lock.Lock()
	defer c.Lock.Unlock()

	var empty T

	if !c.keyExists(key) {
		return empty, false
	}

	// Checks if the entry is expired
	if c.Items[key].IsCacheEntryExpired() {
		c.Lock.Unlock()
		c.RemoveItem(key)
		return empty, false
	}

	return c.Items[key].Item, true
}

// Removes an item
func (c *CacheManager[T]) RemoveItem(key string) {

	c.Lock.Lock()
	defer c.Lock.Unlock()

	if !c.keyExists(key) {
		return
	}

	delete(c.Items, key)
}
