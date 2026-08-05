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

func CreateNewHandler(o *caching.CacheManager[orderstructs.Order]) *OrderHandler {
	return &OrderHandler{
		OrderCache: o,
	}
}

func (oh *OrderHandler) insertItem(order orderstructs.Order, tenantId uint) {
	key := OrderCacheKey + strconv.Itoa(int(order.OrderId))
	oh.OrderCache.SetOrUpdateItem(key, order, tenantId)
}
