package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/internal/customer/helper"
	"ecomps.boobles.cloud/backend/internal/middleware"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels getting a customer by id
func (ch *CustomerHandler) HandleGettingCustomerById(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Customer | HandleGettingCustomerById")

	customerId, err := httputils.IntPathParam(r, "customer_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	key := CustomerCacheKey + strconv.Itoa(customerId)
	cacheItem, ok := ch.CustomerCache.GetItem(key)

	if ok {
		if jsonutils.RespondWithJson(w, http.StatusOK, cacheItem) {
			return
		}
	}

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), ch.Dh, "SelectTenantById", tenantId)

	if !ok {
		fail(http.StatusBadRequest, errors.New("Failed getting tenant"))
		return
	}

	customer, ok := helper.GetCustomer(uint(customerId), tenant.GetPw(ch.Dh, r.Context()), r.Context(), ch.Dh)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting customer"))
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, customer) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (ch *CustomerHandler) HandleGettingAllCustomerByTenantId(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Customer | HandleGettingAllCustomerByTenantId")

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	cacheItems, ok := ch.CustomerCache.GetItems(uint(tenantId))

	if ok {
		if jsonutils.RespondWithJson(w, http.StatusOK, cacheItems) {
			return
		}
	}

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), ch.Dh, "SelectTenantById", tenantId)

	if !ok {
		fail(http.StatusBadRequest, errors.New("Failed getting tenant"))
		return
	}

	allCustomer, ok := helper.GetAllCustomerForTenant(uint(tenantId), tenant.GetPw(ch.Dh, r.Context()), r.Context(), ch.Dh)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting customer"))
		return
	}

	go ch.insertItems(allCustomer)

	if !jsonutils.RespondWithJson(w, http.StatusOK, allCustomer) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
