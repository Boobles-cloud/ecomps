package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"ecomps.boobles.cloud/backend/internal/middleware"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handles getting a order by Id and all its products
func (o *OrderHandler) HandleGettingOrderById(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Order | HandleGettingOrderById")

	orderId, err := httputils.IntPathParam(r, "order_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// Check if the item is in cache
	key := OrderCacheKey + strconv.Itoa(orderId)
	cacheItem, ok := o.orderCache.GetItem(key)

	if ok {
		if jsonutils.RespondWithJson(w, http.StatusOK, cacheItem) {
			return
		}
	}

	order, err := o.orderService.GetOrderById(ctx, uint(orderId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, order) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handels getting all orders for a tenant
func (o *OrderHandler) HandleGettingAllOrdersByTenantId(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Order | HandleGettingAllOrdersByTenantId")

	tenantId := ctx.Value(middleware.TenantIdContextKey).(int)

	cacheItems, ok := o.orderCache.GetItems(uint(tenantId))

	if ok {
		if jsonutils.RespondWithJson(w, http.StatusOK, cacheItems) {
			return
		}
	}

	allOrders, err := o.orderService.GetAllOrdersByTenantId(ctx, uint(tenantId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, allOrders) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handles getting the order status by status_id and language_id
func (o *OrderHandler) HandleGettingStatusById(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Order | HandleGettingStatusById")

	statusId, err := httputils.IntPathParam(r, "status_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	langId, err := httputils.IntPathParam(r, "language_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	status, err := o.statusService.GetById(ctx, uint(statusId), uint(langId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, status) {
		w.WriteHeader(http.StatusOK)
	}
}

// Handles getting all order status by language id
func (o *OrderHandler) HandleGettingAllStatusByLangId(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Order | HandleGettingAllStatusByLangId")

	langId, err := httputils.IntPathParam(r, "language_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	orderStatus, err := o.statusService.GetAllByLangId(ctx, uint(langId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, orderStatus) {
		w.WriteHeader(http.StatusOK)
	}
}
