package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"ecomps.boobles.cloud/backend/internal/middleware"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels getting a customer by id
func (c *CustomerHandler) HandleGettingCustomerById(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Customer | HandleGettingCustomerById")

	customerId, err := httputils.IntPathParam(r, "customer_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	key := CustomerCacheKey + strconv.Itoa(customerId)
	cacheItem, ok := c.customerCache.GetItem(key)

	if ok {
		if jsonutils.RespondWithJson(w, http.StatusOK, cacheItem) {
			return
		}
	}

	customer, err := c.customerService.GetById(ctx, uint(customerId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, customer) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *CustomerHandler) HandleGettingAllCustomerByTenantId(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Customer | HandleGettingAllCustomerByTenantId")

	tenantId := ctx.Value(middleware.TenantIdContextKey).(int)

	cacheItems, ok := c.customerCache.GetItems(uint(tenantId))

	if ok {
		if jsonutils.RespondWithJson(w, http.StatusOK, cacheItems) {
			return
		}
	}

	allCustomer, err := c.customerService.GetAllByTenantId(ctx, uint(tenantId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	go c.insertItems(allCustomer)

	if !jsonutils.RespondWithJson(w, http.StatusOK, allCustomer) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
