package handlers

import (
	"errors"
	"net/http"
	"time"

	"boobles.cloud/backend/database"
	"boobles.cloud/backend/internal/middleware"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
	userstructs "boobles.cloud/backend/internal/user/user_structs"
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

	tenantDeletion, isThere := database.QueryOne[tenantstructs.TenantDeletionStruct](r.Context(), u.Dh, "SelectTenantDeletionFromTenantId", []any{tenantId})

	if tenant.IsUserAdmin(uint(userId)) && !isThere {
		w.Write([]byte("User is still admin in tenant and tenant isn´t deleted"))
		w.WriteHeader(http.StatusConflict)
	}

	if !tenant.IsUserAdmin(uint(userId)) {
		u.Dh.ExecuteSQLStatement("DeleteUserById", []any{userId})
		w.WriteHeader(http.StatusOK)
	}

	userDeletion := userstructs.UserDeletionDatabase{
		IssuedOn:       time.Now(),
		WhenToComplete: tenantDeletion.WhenToComplete,
		UserId:         uint(userId),
	}

	if result := u.Dh.ExecuteSQLStatement("InsertUserDeletion", []any{userDeletion.IssuedOn, userDeletion.WhenToComplete, userDeletion.UserId}); !result.Ok {
		fail(http.StatusInternalServerError, errors.New("Failed creating in database"))
	}

	w.WriteHeader(http.StatusOK)
}
