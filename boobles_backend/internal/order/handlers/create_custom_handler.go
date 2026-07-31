package handlers

import (
	"boobles.cloud/backend/caching"
	orderstructs "boobles.cloud/backend/internal/order/order_structs"
)

type OrderHandler struct {
	OrderCache *caching.CacheManager[orderstructs.Order]
	// ProductCache *caching.CacheManager[] // So we can get products from cache
}

func CreateNewHandler(o *caching.CacheManager[orderstructs.Order]) *OrderHandler {
	return &OrderHandler{
		OrderCache: o,
	}
}
