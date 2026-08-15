package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"boobles.cloud/backend/database"
	"boobles.cloud/backend/internal/middleware"
	orderstructs "boobles.cloud/backend/internal/order/order_structs"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	"boobles.cloud/backend/logging"
)

// Handels creating a order and all its order products in database
func (ho *OrderHandler) HandleCreatingOrder(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(status)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	var order orderstructs.Order

	if err := json.Unmarshal(body, &order); err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// Get the tenant id so we can get the tenant and the password
	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), ho.Dh, "SelectTenantById", tenantId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting tenant"))
		return
	}

	// Create the order
	id, ok := order.CreateOrderInDatabase(tenant.GetPw(ho.Dh, r.Context()), ho.Dh)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to create order"))
		return
	}

	// Loop over the tmp product struct and insert it
	for i := range order.Products {
		if !order.Products[i].InsertIntoDatabase(id, ho.Dh) {
			fail(http.StatusInternalServerError, errors.New("Failed to create product order"))
			return
		}
	}

	order.OrderId = id

	go ho.insertItem(order, uint(tenantId))

	w.WriteHeader(http.StatusOK)
}
