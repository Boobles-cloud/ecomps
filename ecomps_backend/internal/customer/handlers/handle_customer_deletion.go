package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"ecomps.boobles.cloud/backend/internal/middleware"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handles the deletion of a customer
func (c *CustomerHandler) HandleCustomerDeletion(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Customer | HandleCustomerDeletion")

	customerId, err := httputils.IntPathParam(r, "customer_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if err := c.customerService.Delete(ctx, uint(customerId), uint(ctx.Value(middleware.TenantIdContextKey).(int))); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	key := CustomerCacheKey + strconv.Itoa(customerId)
	c.customerCache.RemoveItem(key)
	w.WriteHeader(http.StatusOK)
}
