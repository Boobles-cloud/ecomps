package handlers

import (
	"errors"
	"net/http"

	"boobles.cloud/backend/database"
	"boobles.cloud/backend/internal/middleware"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	"boobles.cloud/backend/logging"
)

// Handle deleting a user
// First check if the user is a admin of the tenant -> he needs to delete the tenant first
// If he is admin -> transfare to new user id or add user to a deletion database and check with every tenant deletion
func (u *UserHandler) HandleUserDeletion(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[User | HandleUserDeletion] "+err.Error())
		w.WriteHeader(status)
	}

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)
	userId := r.Context().Value(middleware.UserIdContextKey).(int)

	tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), u.Dh, "SelectTenantById", []any{tenantId})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting tenant"))
	}

	if tenant.IsUserAdmin(uint(userId)) {
		// TODO: write error msg
		w.WriteHeader(http.StatusConflict)
	}

}
