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

// Handles changing a order
func (o *OrderHandler) HandleChangingOrder(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Order | HandleChangingOrder")

	order, err := jsonutils.JsonDeserilizeHttpRequestBody[orderstructs.Order](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	order.OrderLastChanged = time.Now()

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	if err := o.orderService.UpdateOrder(ctx, order); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	go o.insertItem(order, uint(tenantId))
	w.WriteHeader(http.StatusOK)
}
