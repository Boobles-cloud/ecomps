package handlers

import (
	"strconv"

	customerstructs "ecomps.boobles.cloud/backend/internal/customer/customer_structs"
	"ecomps.boobles.cloud/backend/internal/customer/services"
	"ecomps.boobles.cloud/backend/utils/caching"
)

const (
	CustomerCacheKey = "CUSTOMER:"
)

type CustomerHandler struct {
	customerCache   *caching.CacheManager[customerstructs.Customer]
	customerService *services.CustomerService
}

// Creates a new handler for products
func CreateNewCustomerHandler(c *caching.CacheManager[customerstructs.Customer], cs *services.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerCache:   c,
		customerService: cs,
	}
}

func CustomerToArgs(c customerstructs.Customer) []any {
	return []any{
		c.CustomerId,
		c.CustomerName,
		c.CustomerPostalCode,
		c.CustomerStreetAndHouseNr,
		c.CustomerCity,
		c.CustomerLastChanged,
		c.TenantId,
	}
}

// Use this func to set all cache items
// This func is used in a seperate go routine
// We just fire and forgett about it, because we can live without a cache
func (c *CustomerHandler) insertItems(t []customerstructs.Customer) {

	for i := range t {
		key := CustomerCacheKey + strconv.Itoa(int(t[i].CustomerId))
		c.customerCache.SetOrUpdateItem(key, t[i], t[i].TenantId)
	}
}

// Use this func to set all cache items
// This func is used in a seperate go routine
// We just fire and forgett about it, because we can live without a cache
func (c *CustomerHandler) insertItem(t customerstructs.Customer) {
	key := CustomerCacheKey + strconv.Itoa(int(t.CustomerId))
	c.customerCache.SetOrUpdateItem(key, t, t.TenantId)
}
