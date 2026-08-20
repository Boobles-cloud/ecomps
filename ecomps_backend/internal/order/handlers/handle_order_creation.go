package handlers

import (
	"errors"
	"net/http"

	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/internal/middleware"
	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels creating a order and all its order products in database
func (ho *OrderHandler) HandleCreatingOrder(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Order | HandleCreatingOrder")

	order, err := jsonutils.JsonDeserilizeHttpRequestBody[orderstructs.Order](r)

	if err != nil {
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
