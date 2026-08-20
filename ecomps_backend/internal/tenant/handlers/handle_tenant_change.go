package handlers

import (
	"net/http"

	"ecomps.boobles.cloud/backend/database"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels a tenant change
func (t *TenantHandler) HandleTenantChange(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Tenant | HandleTenantChange")

	// TODO: Check why we need this here?
	wantedUpdateType := r.URL.Query().Get("type")

	if wantedUpdateType == "" {
		fail(http.StatusBadRequest, nil)
		return
	}

	tenant, err := jsonutils.JsonDeserilizeHttpRequestBody[tenantstructs.Tenant](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if !database.UpdateDatabaseEntry[tenantstructs.Tenant](t.Dh, "UpdateTenant", "TenantId", tenant) {
		fail(http.StatusInternalServerError, nil)
		return
	}

	w.WriteHeader(http.StatusOK)
}
