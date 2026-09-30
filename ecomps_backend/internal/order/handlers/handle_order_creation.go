package handlers

import (
	"context"
	"net/http"
	"time"

	"ecomps.boobles.cloud/backend/internal/middleware"
	orderstructs "ecomps.boobles.cloud/backend/internal/order/order_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels creating a order and all its order products in database
func (o *OrderHandler) HandleCreatingOrder(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Order | HandleCreatingOrder")

	order, err := jsonutils.JsonDeserilizeHttpRequestBody[orderstructs.Order](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// Get the tenant id so we can get the tenant and the password
	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	order.TenantId = uint(tenantId)

	id, err := o.orderService.CreateOrder(ctx, order)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	order.OrderId = id

	go o.insertItem(order, uint(tenantId))

	w.WriteHeader(http.StatusOK)
}
