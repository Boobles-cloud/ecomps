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

	cookie, err := httputils.CreateAuthCookie(r.Context().Value(middleware.UserIdContextKey).(uint), tenantStruct.TenantId, t.Dh)

	if err != nil {
		// TODO: Change this here, so the frontend nows that the user needs to be logged out again
		// This will be added with the whole refactoring of the error stuff
		jsonutils.RespondWithJson(w, http.StatusOK, tenantStruct)
		return
	}

	http.SetCookie(w, &cookie)

	if !jsonutils.RespondWithJson(w, http.StatusOK, tenantStruct) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
