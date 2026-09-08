package handlers

import (
	"context"
	"net/http"
	"time"

	"ecomps.boobles.cloud/backend/internal/middleware"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// ============= NOTE =============
// 	This needs the frontend auth
// ================================

// Handles getting the user by the user id
func (hu *UserHandler) HandleGettingUserById(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "User | HandleGettingUserById")

	userId, err := httputils.IntPathParam(r, "user_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	user, err := hu.userService.GetUserById(ctx, uint(userId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, user) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handles the request for checking if a user has a tenant
func (hu *UserHandler) HandleHasUserATenant(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "User | HandleHasUserATenant")

	userId, err := httputils.IntPathParam(r, "user_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	user, err := hu.userService.GetUserById(ctx, uint(userId))

	if user.UserHasTenant && user.TenantId != 0 {
		if !jsonutils.RespondWithJson(w, http.StatusOK, "Tenant = true") {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, "Tenant = false") {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Handels getting all users for a tenant
// The tenant id is providet by our auth middleware
func (hu *UserHandler) HandleGettingAllUsersByUserTenantId(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Tenant | HandleGettingAllUsersByUserTenantId")

	tenantId := r.Context().Value(middleware.TenantIdContextKey).(int)

	users, err := hu.userService.GetAllByTenantId(ctx, uint(tenantId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, users) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
