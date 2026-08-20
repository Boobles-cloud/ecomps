package handlers

import (
	"errors"
	"net/http"

	"ecomps.boobles.cloud/backend/database"
	"ecomps.boobles.cloud/backend/internal/middleware"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels the request for getting a tenant by the user Id
func (t *TenantHandler) HandleGetTenantByUserId(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Tenant | HandleGetTenantByUserId")

	userIdInt, err := httputils.IntPathParam(r, "user_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), t.Dh, "SelectTenantByUserId", userIdInt)

	if !ok {
		fail(http.StatusBadRequest, nil)
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, tenant) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handels getting the tenant by the given id
func (t *TenantHandler) HandleGetTenantByTenantId(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Tenant | HandleGetTenantByTenantId")

	tenantId, err := httputils.IntPathParam(r, "tenant_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), t.Dh, "SelectTenantById", tenantId)

	if !ok {
		fail(http.StatusBadRequest, errors.New("Failed getting tenant"))
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, tenant) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handels getting all users for a tenant
// The tenant id is providet by our auth middleware
func (t *TenantHandler) HandleGettingAllUsersByUserTenantId(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Tenant | HandleGettingAllUsersByUserTenantId")

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	users, ok := database.QueryMany[userstructs.UserStruct](r.Context(), t.Dh, "SelectAllUsersByTenant", tenantId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting users for tenant"))
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, users) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
