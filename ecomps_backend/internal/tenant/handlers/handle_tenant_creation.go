package handlers

import (
	"net/http"

	"ecomps.boobles.cloud/backend/internal/middleware"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handles the tenant creation
func (t *TenantHandler) HandleTenantCreation(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Tenant | HandleTenantCreation")

	// Read the body
	tenantStruct, err := jsonutils.JsonDeserilizeHttpRequestBody[tenantstructs.Tenant](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// Create the tenant
	if !tenantStruct.CreateTenantInDatabase(r.Context(), r.Context().Value(middleware.UserIdContextKey).(int), t.Dh) {
		fail(http.StatusInternalServerError, nil)
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, tenantStruct) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
