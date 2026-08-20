package handlers

import (
	"net/http"
	"time"

	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels the deletion of a tenant
func (t *TenantHandler) HandleTenantDeletion(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Tenant | HandleTenantDeletion")

	tenantDeleteStruct, err := jsonutils.JsonDeserilizeHttpRequestBody[tenantstructs.TenantDeletionStruct](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// We set some vars here
	tenantDeleteStruct.IssuedOn = time.Now()
	tenantDeleteStruct.WhenToComplete = time.Now().AddDate(0, 2, 0)
	tenantDeleteStruct.Deleted = false

	if result := t.Dh.ExecuteSQLStatement("InsertTenantDeletion", []any{tenantDeleteStruct.IssuedFrom, tenantDeleteStruct.IssuedOn, tenantDeleteStruct.WhenToComplete, tenantDeleteStruct.Deleted, tenantDeleteStruct.TenantId}); !result.Ok {
		fail(http.StatusInternalServerError, nil)
		return
	}

	w.WriteHeader(http.StatusOK)
}
