package handlers

import (
	"strconv"

	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
	"ecomps.boobles.cloud/backend/internal/order/services"
	"ecomps.boobles.cloud/backend/utils/caching"
)

const (
	OrderCacheKey = "ORDER:"
)

type OrderHandler struct {
	orderCache    *caching.CacheManager[orderstructs.Order]
	orderService  *services.OrderService
	statusService *services.StatusService
}

func CreateNewOrderHandler(oc *caching.CacheManager[orderstructs.Order], o *services.OrderService, s *services.StatusService) *OrderHandler {
	return &OrderHandler{
		orderCache:    oc,
		orderService:  o,
		statusService: s,
	}
}

// Use this func to set all cache items
// This func is used in a seperate go routine
// We just fire and forgett about it, because we can live without a cache
func (o *OrderHandler) insertItems(t []orderstructs.Order, tenantId uint) {

	for i := range t {

		key := OrderCacheKey + strconv.Itoa(int(t[i].OrderId))
		o.orderCache.SetOrUpdateItem(key, t[i], tenantId)
	}
}

// Use this func to set all cache items
// This func is used in a seperate go routine
// We just fire and forgett about it, because we can live without a cache
func (o *OrderHandler) insertItem(order orderstructs.Order, tenantId uint) {
	key := OrderCacheKey + strconv.Itoa(int(order.OrderId))
	o.orderCache.SetOrUpdateItem(key, order, tenantId)
}
