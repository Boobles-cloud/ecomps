package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"boobles.cloud/backend/database"
	"boobles.cloud/backend/internal/middleware"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	userstructs "boobles.cloud/backend/internal/user/user_structs"
	"boobles.cloud/backend/logging"
)

// Handels the request for getting a tenant by the user Id
func (t *TenantHandler) HandleGetTenantByUserId(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Tenant | HandleGetTenantByUserId] "+err.Error())
		}

		w.WriteHeader(status)
	}

	userIdInt, err := strconv.Atoi(r.PathValue("user-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), t.Dh, "SelectTenantByUserId", []any{userIdInt})

	if !ok {
		fail(http.StatusBadRequest, nil)
	}

	tenantJson, err := json.Marshal(tenant)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(tenantJson)
	w.WriteHeader(http.StatusOK)
}

// Handels getting the tenant by the given id
func (t *TenantHandler) HandleGetTenantByTenantId(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Tenant | HandleGetTenantByUserId] "+err.Error())

		w.WriteHeader(status)
	}

	tenantId, err := strconv.Atoi(r.PathValue("tenant-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), t.Dh, "SelectTenantById", []any{tenantId})

	if !ok {
		fail(http.StatusBadRequest, errors.New("Failed getting tenant"))
	}

	tenantJson, err := json.Marshal(tenant)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(tenantJson)
	w.WriteHeader(http.StatusOK)
}

// Handels getting all users for a tenant
// The tenant id is providet by our auth middleware
func (t *TenantHandler) HandleGettingAllUsersByUserTenantId(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Tenant | HandleGettingAllUsersByTenantId] "+err.Error())
		w.WriteHeader(status)
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	users, ok := database.QueryMany[userstructs.UserStruct](r.Context(), t.Dh, "SelectAllUsersByTenant", []any{tenantId})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting users for tenant"))
	}

	jsonData, err := json.Marshal(users)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}
