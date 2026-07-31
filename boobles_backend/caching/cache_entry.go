package caching

import (
	"time"
)

type CacheEntry[T any] struct {
	Item           T
	ExpirationTime time.Time
}

// Creates a new cache entry
func CreateNewCacheEntry[T any](item T) CacheEntry[T] {
	return CacheEntry[T]{
		Item:           item,
		ExpirationTime: time.Now().Add(2 * time.Hour),
	}
}

// Checks if the cache entry is expired
func (c CacheEntry[T]) IsCacheEntryExpired() bool {

	if c.ExpirationTime.After(time.Now()) {
		return true
	}

	return false
}
