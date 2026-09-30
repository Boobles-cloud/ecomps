package handlers

import (
	"context"
	"net/http"
	"time"

	"ecomps.boobles.cloud/backend/internal/middleware"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handle deleting a user
// First check if the user is a admin of the tenant -> he needs to delete the tenant first
// If he is admin -> transfare to new user id or add user to a deletion database and check with every tenant deletion
func (hu *UserHandler) HandleUserDeletion(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "User | HandleUserDeletion")

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)
	userId := r.Context().Value(middleware.UserIdContextKey).(int)

	if err := hu.userService.DeleteUser(ctx, uint(userId), uint(tenantId)); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
