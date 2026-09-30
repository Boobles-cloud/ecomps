package handlers

import (
	"context"
	"net/http"
	"time"

	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels getting the tenant by the given id
func (t *TenantHandler) HandleGetTenantByTenantId(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Tenant | HandleGetTenantByTenantId")

	tenantId, err := httputils.IntPathParam(r, "tenant_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenant, err := t.tenantService.GetTenantById(ctx, uint(tenantId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, tenant) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
