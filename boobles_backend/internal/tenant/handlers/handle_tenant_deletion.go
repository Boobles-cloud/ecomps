package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"boobles.cloud/backend/database"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	"boobles.cloud/backend/logging"
)

// TODO: Come back here, when the permissions stuff is ready!!!

// Handels the deletion of a tenant
func HandleTenantDeltion(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Tenant | HandleTenantDeletion] "+err.Error())
		}

		w.WriteHeader(status)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	var tenantDeleteStruct tenantstructs.TenantDeletionStruct
	if err := json.Unmarshal(body, &tenantDeleteStruct); err != nil {
		fail(http.StatusBadRequest, err)
	}

	// We set some vars here
	tenantDeleteStruct.IssuedOn = time.Now()
	tenantDeleteStruct.WhenToComplete = time.Now().AddDate(0, 2, 0)
	tenantDeleteStruct.Deleted = false

	if result := database.ExecuteSQLStatement("InsertTenantDeletion", database.Insert, []any{tenantDeleteStruct.IssuedFrom, tenantDeleteStruct.IssuedOn, tenantDeleteStruct.WhenToComplete, tenantDeleteStruct.Deleted, tenantDeleteStruct.TenantId}); !result.Ok {
		fail(http.StatusInternalServerError, nil)
	}

	w.WriteHeader(http.StatusOK)
}
