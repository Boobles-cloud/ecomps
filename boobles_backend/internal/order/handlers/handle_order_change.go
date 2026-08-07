package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"boobles.cloud/backend/crypto"
	"boobles.cloud/backend/database"
	"boobles.cloud/backend/internal/middleware"
	orderstructs "boobles.cloud/backend/internal/order/order_structs"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	"boobles.cloud/backend/logging"
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
	}

	var order orderstructs.Order

	if err := json.Unmarshal(body, &order); err != nil {
		fail(http.StatusBadRequest, err)
	}

	order.OrderLastChanged = time.Now()

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)
	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), ho.Dh, "SelectTenantById", []any{tenantId})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get tenant"))
	}

	encryptedOrder, ok := crypto.Encrypt[orderstructs.Order](order, tenant.GetPw(ho.Dh, r.Context()))

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed encrypting order"))
	}

	if !database.UpdateDatabaseEntry[orderstructs.Order](ho.Dh, "UpdateOrder", "OrderId", encryptedOrder) {
		fail(http.StatusInternalServerError, errors.New("Failed updating order"))
	}

	for i := range order.Products {
		order.Products[i].UpdateOrderProduct(ho.Dh)
	}

	go ho.insertItem(order, uint(tenantId))
	w.WriteHeader(http.StatusOK)
}
