package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handles deleting a order
func (o OrderHandler) HandleOrderDeletion(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Order | HandleOrderDeletion")

	orderId, err := httputils.IntPathParam(r, "order_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if err := o.orderService.DeleteOrder(ctx, uint(orderId)); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	key := OrderCacheKey + strconv.Itoa(orderId)
	o.orderCache.RemoveItem(key)

	w.WriteHeader(http.StatusOK)
}
