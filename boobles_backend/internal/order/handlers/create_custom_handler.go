package handlers

import (
	"strconv"

	"boobles.cloud/backend/caching"
	orderstructs "boobles.cloud/backend/internal/order/order_structs"
)

const (
	OrderCacheKey = "ORDER:"
)

type OrderHandler struct {
	OrderCache *caching.CacheManager[orderstructs.Order]
}

func CreateNewOrderHandler(o *caching.CacheManager[orderstructs.Order]) *OrderHandler {
	return &OrderHandler{
		OrderCache: o,
	}
}

// Use this func to set all cache items
// This func is used in a seperate go routine
// We just fire and forgett about it, because we can live without a cache
func (o *OrderHandler) insertItems(t []orderstructs.Order, tenantId uint) {

	for i := range t {

		key := OrderCacheKey + strconv.Itoa(int(t[i].OrderId))
		o.OrderCache.SetOrUpdateItem(key, t[i], tenantId)
	}
}

// Use this func to set all cache items
// This func is used in a seperate go routine
// We just fire and forgett about it, because we can live without a cache
func (oh *OrderHandler) insertItem(order orderstructs.Order, tenantId uint) {
	key := OrderCacheKey + strconv.Itoa(int(order.OrderId))
	oh.OrderCache.SetOrUpdateItem(key, order, tenantId)
}
