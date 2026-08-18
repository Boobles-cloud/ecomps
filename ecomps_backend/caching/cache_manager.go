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
func (c *CacheManager[T]) SetOrUpdateItem(key string, item T, tenantId uint) {

	c.Lock.Lock()
	defer c.Lock.Unlock()

	entry := CreateNewCacheEntry[T](item, tenantId)
	c.Items[key] = entry
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

// Gets all items for a tenantId
func (c *CacheManager[T]) GetItems(tenantId uint) ([]T, bool) {

	c.Lock.Lock()
	defer c.Lock.Unlock()

	allItems := make([]T, 100)

	for k := range c.Items {
		if c.Items[k].TenantId == tenantId {
			allItems = append(allItems, c.Items[k].Item)
		}
	}

	return allItems, len(allItems) != 0
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

// Checks if a key exists
func (c *CacheManager[T]) keyExists(key string) bool {
	for k := range c.Items {
		if k == key {
			return true
		}
	}
	return false
}
