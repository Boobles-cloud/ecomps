package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"boobles.cloud/backend/database"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	"boobles.cloud/backend/logging"
)

// Handels the request for getting a tenant by the user Id
func HandleGetTenantByUserId(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Tenant | HandleGetTenantByUserId] "+err.Error())
		}

		w.WriteHeader(status)
	}

	userIdInt, err := strconv.Atoi(r.PathValue("user-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	tenant, ok := database.QueryDatabase[tenantstructs.Tenant]("SelectTenantByUserId", []any{userIdInt})

	if !ok || len(tenant) > 1 {
		fail(http.StatusBadRequest, nil)
	}

	tenantJson, err := json.Marshal(tenant[0])

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(tenantJson)
	w.WriteHeader(http.StatusOK)
}

// Handels getting the tenant by the given id
func HandleGetTenantByTenantId(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Tenant | HandleGetTenantByUserId] "+err.Error())
		}

		w.WriteHeader(status)
	}

	tenantId, err := strconv.Atoi(r.PathValue("tenant-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	tenant, ok := database.QueryDatabase[tenantstructs.Tenant]("SelectTenantById", []any{tenantId})

	if !ok || len(tenant) > 1 {
		fail(http.StatusBadRequest, nil)
	}

	tenantJson, err := json.Marshal(tenant[0])

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(tenantJson)
	w.WriteHeader(http.StatusOK)
}
