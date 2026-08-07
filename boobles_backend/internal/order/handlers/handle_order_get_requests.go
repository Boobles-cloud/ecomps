package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"boobles.cloud/backend/database"
	"boobles.cloud/backend/internal/middleware"
	"boobles.cloud/backend/internal/order/helper"
	orderstructs "boobles.cloud/backend/internal/order/order_structs"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	"boobles.cloud/backend/logging"
)

// Handles getting a order by Id and all its products
func (ho *OrderHandler) HandleGettingOrderById(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Order | HandleGettingOrderById] "+err.Error())
		w.WriteHeader(status)
	}

	orderId, err := strconv.Atoi(r.PathValue("order-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
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

		w.Write(jsonData)
		w.WriteHeader(http.StatusOK)
	}

withOutCache:
	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), ho.Dh, "SelectTenantById", []any{tenantId})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting tenant"))
	}

	// Gets the encrypted order
	order, ok := helper.GetOrder(uint(orderId), tenant.GetPw(ho.Dh, r.Context()), ho.Dh, r.Context())

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting order"))
	}

	// Get all products for this order
	order.GetAllProducts(r.Context(), ho.Dh)

	jsonData, err := json.Marshal(*order)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	// Write it into cache
	go ho.insertItem(*order, uint(tenantId))

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
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

		w.Write(jsonData)
		w.WriteHeader(http.StatusOK)
	}

withOutCache:

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), ho.Dh, "SelectTenantById", []any{tenantId})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting tenant"))
	}

	allOrders, ok := helper.GetAllOrders(uint(tenantId), tenant.GetPw(ho.Dh, r.Context()), ho.Dh, r.Context())

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting orders"))
	}

	jsonData, err := json.Marshal(allOrders)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	go ho.insertItems(allOrders, uint(tenantId))

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

// Handles getting the order status by status-id and language-id
func (ho *OrderHandler) HandleGettingStatusById(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Order | HandleGettingStatusById] "+err.Error())
		w.WriteHeader(status)
	}

	statusId, err := strconv.Atoi(r.PathValue("status-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	langId, err := strconv.Atoi(r.PathValue("language-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	status, ok := database.QueryOne[orderstructs.OrderStatus](r.Context(), ho.Dh, "SelectOrderStatusByIdAndLanguageId", []any{statusId, langId})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting status"))
	}

	jsonData, err := json.Marshal(status)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

// Handles getting all order status by language id
func (ho *OrderHandler) HandleGettingAllStatusByLangId(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Order | HandeGettingallStatusByLangId]"+err.Error())
		w.WriteHeader(status)
	}

	langId, err := strconv.Atoi(r.PathValue("language-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	orderStatus, ok := database.QueryMany[orderstructs.OrderStatus](r.Context(), ho.Dh, "SelectAllStatusByLanguageId", []any{langId})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting all order status"))
	}

	jsonData, err := json.Marshal(orderStatus)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}
