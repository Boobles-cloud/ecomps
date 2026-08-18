package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/internal/middleware"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	"ecomps.boobles.cloud/backend/utils/logging"
)

// Handels the request for getting a tenant by the user Id
func (t *TenantHandler) HandleGetTenantByUserId(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Tenant | HandleGetTenantByUserId] "+err.Error())
		}

		w.WriteHeader(status)
	}

	userIdInt, err := strconv.Atoi(r.PathValue("user_id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), t.Dh, "SelectTenantByUserId", userIdInt)

	if !ok {
		fail(http.StatusBadRequest, nil)
		return
	}

	tenantJson, err := json.Marshal(tenant)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tenantJson)
}

// Handels getting the tenant by the given id
func (t *TenantHandler) HandleGetTenantByTenantId(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Tenant | HandleGetTenantByUserId] "+err.Error())

		w.WriteHeader(status)
	}

	tenantId, err := strconv.Atoi(r.PathValue("tenant_id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), t.Dh, "SelectTenantById", tenantId)

	if !ok {
		fail(http.StatusBadRequest, errors.New("Failed getting tenant"))
		return
	}

	tenantJson, err := json.Marshal(tenant)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tenantJson)
}

// Handels getting all users for a tenant
// The tenant id is providet by our auth middleware
func (t *TenantHandler) HandleGettingAllUsersByUserTenantId(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Tenant | HandleGettingAllUsersByTenantId] "+err.Error())
		w.WriteHeader(status)
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	users, ok := database.QueryMany[userstructs.UserStruct](r.Context(), t.Dh, "SelectAllUsersByTenant", tenantId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting users for tenant"))
		return
	}

	jsonData, err := json.Marshal(users)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
