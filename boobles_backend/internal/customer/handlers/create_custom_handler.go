package handlers

import (
	"strconv"

	"boobles.cloud/backend/caching"
	"boobles.cloud/backend/database"
	customerstructs "boobles.cloud/backend/internal/customer/customer_structs"
)

const (
	CustomerCacheKey = "CUSTOMER:"
)

type CustomerHandler struct {
	CustomerCache *caching.CacheManager[customerstructs.Customer]
	Dh            *database.DbHandler
}

// Creates a new handler for products
func CreateNewCustomerHandler(c *caching.CacheManager[customerstructs.Customer], d *database.DbHandler) *CustomerHandler {
	return &CustomerHandler{
		CustomerCache: c,
		Dh:            d,
	}
}

// Use this func to set all cache items
// This func is used in a seperate go routine
// We just fire and forgett about it, because we can live without a cache
func (p *CustomerHandler) insertItems(t []customerstructs.Customer) {

	for i := range t {
		key := CustomerCacheKey + strconv.Itoa(int(t[i].CustomerId))
		p.CustomerCache.SetOrUpdateItem(key, t[i], t[i].TenantId)
	}
}

// Use this func to set all cache items
// This func is used in a seperate go routine
// We just fire and forgett about it, because we can live without a cache
func (p *CustomerHandler) insertItem(t customerstructs.Customer) {
	key := CustomerCacheKey + strconv.Itoa(int(t.CustomerId))
	p.CustomerCache.SetOrUpdateItem(key, t, t.TenantId)
}
