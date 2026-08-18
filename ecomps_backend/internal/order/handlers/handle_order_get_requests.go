package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/internal/middleware"
	"ecomps.boobles.cloud/backend/internal/order/helper"
	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	"ecomps.boobles.cloud/backend/utils/logging"
)

// Handles getting a order by Id and all its products
func (ho *OrderHandler) HandleGettingOrderById(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Order | HandleGettingOrderById] "+err.Error())
		w.WriteHeader(status)
	}

	orderId, err := strconv.Atoi(r.PathValue("order_id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// Check if the item is in cache
	key := OrderCacheKey + strconv.Itoa(orderId)
	cacheItem, ok := ho.OrderCache.GetItem(key)

	if ok {

		jsonData, err := json.Marshal(cacheItem)

		if err != nil {
			// If there is an error, just continue there
			goto withOutCache
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}

withOutCache:
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

	jsonData, err := json.Marshal(*order)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	// Write it into cache
	go ho.insertItem(*order, uint(tenantId))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Handels getting all orders for a tenant
func (ho *OrderHandler) HandleGettingAllOrdersByTenantId(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Order | HandleGettingAllOrdersByTenantId] "+err.Error())
		w.WriteHeader(status)
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	cacheItems, ok := ho.OrderCache.GetItems(uint(tenantId))

	if ok {
		jsonData, err := json.Marshal(cacheItems)

		if err != nil {
			goto withOutCache
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}

withOutCache:

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

	jsonData, err := json.Marshal(allOrders)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	go ho.insertItems(allOrders, uint(tenantId))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Handles getting the order status by status_id and language_id
func (ho *OrderHandler) HandleGettingStatusById(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Order | HandleGettingStatusById] "+err.Error())
		w.WriteHeader(status)
	}

	statusId, err := strconv.Atoi(r.PathValue("status_id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	langId, err := strconv.Atoi(r.PathValue("language_id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	status, ok := database.QueryOne[orderstructs.OrderStatus](r.Context(), ho.Dh, "SelectOrderStatusByIdAndLanguageId", statusId, langId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting status"))
		return
	}

	jsonData, err := json.Marshal(status)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Handles getting all order status by language id
func (ho *OrderHandler) HandleGettingAllStatusByLangId(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Order | HandeGettingallStatusByLangId]"+err.Error())
		w.WriteHeader(status)
	}

	langId, err := strconv.Atoi(r.PathValue("language_id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	orderStatus, ok := database.QueryMany[orderstructs.OrderStatus](r.Context(), ho.Dh, "SelectAllStatusByLanguageId", langId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting all order status"))
		return
	}

	jsonData, err := json.Marshal(orderStatus)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
