package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"ecomps.boobles.cloud/backend/crypto"
	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/internal/middleware"
	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	"ecomps.boobles.cloud/backend/logging"
)

// Handles changing a order
func (ho *OrderHandler) HandleChangingOrder(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Order | HandleChangingOrder] "+err.Error())
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

	order.OrderLastChanged = time.Now()

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)
	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), ho.Dh, "SelectTenantById", tenantId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get tenant"))
		return
	}

	encryptedOrder, ok := crypto.Encrypt[orderstructs.Order](order, tenant.GetPw(ho.Dh, r.Context()))

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed encrypting order"))
		return
	}

	if !database.UpdateDatabaseEntry[orderstructs.Order](ho.Dh, "UpdateOrder", "OrderId", encryptedOrder) {
		fail(http.StatusInternalServerError, errors.New("Failed updating order"))
		return
	}

	for i := range order.Products {
		order.Products[i].UpdateOrderProduct(ho.Dh)
	}

	go ho.insertItem(order, uint(tenantId))
	w.WriteHeader(http.StatusOK)
}
