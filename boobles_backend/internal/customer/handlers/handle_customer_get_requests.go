package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"boobles.cloud/backend/database"
	"boobles.cloud/backend/internal/customer/helper"
	"boobles.cloud/backend/internal/middleware"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	"boobles.cloud/backend/logging"
)

// Handels getting a customer by id
func (ch *CustomerHandler) HandleGettingCustomerById(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Customer | HandleGettingCustomerById] "+err.Error())
		w.WriteHeader(status)
	}

	customerId, err := strconv.Atoi(r.PathValue("customer-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
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

		w.Write(jsonData)
		w.WriteHeader(http.StatusOK)
	}

withOutCache:

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), ch.Dh, "SelectTenantById", []any{tenantId})

	if !ok {
		fail(http.StatusBadRequest, errors.New("Failed getting tenant"))
	}

	customer, ok := helper.GetCustomer(uint(customerId), tenant.GetPw(ch.Dh, r.Context()), r.Context(), ch.Dh)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting customer"))
	}

	go ch.insertItem(customer)

	jsonData, err := json.Marshal(customer)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
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

		w.Write(jsonData)
		w.WriteHeader(http.StatusOK)
	}

withOutCache:

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), ch.Dh, "SelectTenantById", []any{tenantId})

	if !ok {
		fail(http.StatusBadRequest, errors.New("Failed getting tenant"))
	}

	allCustomer, ok := helper.GetAllCustomerForTenant(uint(tenantId), tenant.GetPw(ch.Dh, r.Context()), r.Context(), ch.Dh)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting customer"))
	}

	go ch.insertItems(allCustomer)

	jsonData, err := json.Marshal(allCustomer)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}
