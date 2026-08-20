package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/internal/middleware"
	"ecomps.boobles.cloud/backend/internal/order/helper"
	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handles getting a order by Id and all its products
func (ho *OrderHandler) HandleGettingOrderById(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Order | HandleGettingOrderById")

	orderId, err := httputils.IntPathParam(r, "order_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// Check if the item is in cache
	key := OrderCacheKey + strconv.Itoa(orderId)
	cacheItem, ok := ho.OrderCache.GetItem(key)

	if ok {

		if jsonutils.RespondWithJson(w, http.StatusOK, cacheItem) {
			return
		}
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), ho.Dh, "SelectTenantById", tenantId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting tenant"))
		return
	}

	// Gets the encrypted order
	order, ok := helper.GetOrder(uint(orderId), tenant.GetPw(ho.Dh, r.Context()), ho.Dh, r.Context())

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting order"))
		return
	}

	// Get all products for this order
	order.GetAllProducts(r.Context(), ho.Dh)

	if !jsonutils.RespondWithJson(w, http.StatusOK, order) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handels getting all orders for a tenant
func (ho *OrderHandler) HandleGettingAllOrdersByTenantId(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Order | HandleGettingAllOrdersByTenantId")

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	cacheItems, ok := ho.OrderCache.GetItems(uint(tenantId))

	if ok {
		if jsonutils.RespondWithJson(w, http.StatusOK, cacheItems) {
			return
		}
	}

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), ho.Dh, "SelectTenantById", tenantId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting tenant"))
		return
	}

	allOrders, ok := helper.GetAllOrders(uint(tenantId), tenant.GetPw(ho.Dh, r.Context()), ho.Dh, r.Context())

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting orders"))
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, allOrders) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handles getting the order status by status_id and language_id
func (ho *OrderHandler) HandleGettingStatusById(w http.ResponseWriter, r *http.Request) {

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

	status, ok := database.QueryOne[orderstructs.OrderStatus](r.Context(), ho.Dh, "SelectOrderStatusByIdAndLanguageId", statusId, langId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting status"))
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, status) {
		w.WriteHeader(http.StatusOK)
	}
}

// Handles getting all order status by language id
func (ho *OrderHandler) HandleGettingAllStatusByLangId(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Order | HandleGettingAllStatusByLangId")

	langId, err := httputils.IntPathParam(r, "language_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	orderStatus, ok := database.QueryMany[orderstructs.OrderStatus](r.Context(), ho.Dh, "SelectAllStatusByLanguageId", langId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting all order status"))
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, orderStatus) {
		w.WriteHeader(http.StatusOK)
	}
}
