package handlers

import (
	"context"
	"net/http"
	"time"

	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels a tenant change
func (t *TenantHandler) HandleTenantChange(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Tenant | HandleTenantChange")

	tenant, err := jsonutils.JsonDeserilizeHttpRequestBody[tenantstructs.Tenant](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if err := t.tenantService.UpdateTenant(ctx, tenant); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
