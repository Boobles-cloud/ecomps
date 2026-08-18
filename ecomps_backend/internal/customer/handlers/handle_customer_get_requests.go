package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/internal/customer/helper"
	"ecomps.boobles.cloud/backend/internal/middleware"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	"ecomps.boobles.cloud/backend/utils/logging"
)

// Handels getting a customer by id
func (ch *CustomerHandler) HandleGettingCustomerById(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Customer | HandleGettingCustomerById] "+err.Error())
		w.WriteHeader(status)
	}

	customerId, err := strconv.Atoi(r.PathValue("customer_id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	key := CustomerCacheKey + strconv.Itoa(customerId)
	cacheItem, ok := ch.CustomerCache.GetItem(key)

	if ok {

		jsonData, err := json.Marshal(cacheItem)

		if err != nil {
			// If there is an error, just continue there
			goto withOutCache
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}

withOutCache:

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

	go ch.insertItem(customer)

	jsonData, err := json.Marshal(customer)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (ch *CustomerHandler) HandleGettingAllCustomerByTenantId(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Customer | HandleGettingCustomerById] "+err.Error())
		w.WriteHeader(status)
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	cacheItems, ok := ch.CustomerCache.GetItems(uint(tenantId))

	if ok {

		jsonData, err := json.Marshal(cacheItems)

		if err != nil {
			// If there is an error, just continue there
			goto withOutCache
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}

withOutCache:

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

	jsonData, err := json.Marshal(allCustomer)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
