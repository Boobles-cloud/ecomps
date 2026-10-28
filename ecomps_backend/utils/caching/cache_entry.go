package caching

import (
	"time"
	"uuid"
)

type CacheEntry[T any] struct {
	TenantId       uuid.UUID
	Item           T
	ExpirationTime time.Time
}

// Creates a new cache entry
func CreateNewCacheEntry[T any](item T, tenantId uuid.UUID) CacheEntry[T] {
	return CacheEntry[T]{
		TenantId:       tenantId,
		Item:           item,
		ExpirationTime: time.Now().Add(2 * time.Hour),
	}
}

// Checks if the cache entry is expired
func (c CacheEntry[T]) IsCacheEntryExpired() bool {
	return time.Now().After(c.ExpirationTime)
}
