package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"boobles.cloud/backend/database"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	"boobles.cloud/backend/logging"
)

func (t *TenantHandler) HandleTenantChange(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Tenant | HandleTenantChange] "+err.Error())
		}

		w.WriteHeader(status)
	}

	wantedUpdateType := r.URL.Query().Get("type")

	if wantedUpdateType == "" {
		fail(http.StatusBadRequest, nil)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	var tenant tenantstructs.Tenant
	if err := json.Unmarshal(body, &tenant); err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if !database.UpdateDatabaseEntry[tenantstructs.Tenant](t.Dh, "UpdateTenant", "TenantId", tenant) {
		fail(http.StatusInternalServerError, nil)
		return
	}

	w.WriteHeader(http.StatusOK)
}
