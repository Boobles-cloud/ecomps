package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"boobles.cloud/backend/internal/middleware"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	"boobles.cloud/backend/logging"
)

// Handles the tenant creation
func (t *TenantHandler) HandleTenantCreation(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Tenant | HandleTenantCreatioon] "+err.Error())
		}

		w.WriteHeader(status)
	}

	// Read the body
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	// Unmarshal everything
	var tenantStruct tenantstructs.Tenant
	if err := json.Unmarshal(body, &tenantStruct); err != nil {
		fail(http.StatusBadRequest, err)
	}

	// Create the tenant
	if !tenantStruct.CreateTenantInDatabase(r.Context().Value(middleware.UserIdContextKey).(int), t.Dh) {
		fail(http.StatusInternalServerError, nil)
	}

	// Marshal our Tenant so the frontend gets the tenant Id
	finalTenant, err := json.Marshal(tenantStruct)
	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(finalTenant)
	w.WriteHeader(http.StatusOK)
}
