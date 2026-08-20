package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"ecomps.boobles.cloud/backend/database"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// ============= NOTE =============
// 	This needs the frontend auth
// ================================

// Handles the getting the user by the accesstoken val
func (u *UserHandler) HandleGettingUserByAuthTokenVal(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "User | HanldeGettingUserByAuthTokenVal")

	authToken := r.PathValue("authtoken")

	if authToken == "" {
		fail(http.StatusBadRequest, errors.New("Failed getting authtoken or token is null"))
		return
	}

	user, ok := database.QueryOne[userstructs.UserStruct](r.Context(), u.Dh, "SelectUserByAuthToken", authToken)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting user or user is more then one"))
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, user) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handles getting the user by the user id
func (u *UserHandler) HandleGettingUserById(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "User | HandleGettingUserById")

	userId, err := strconv.Atoi(r.PathValue("user_id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	user, ok := database.QueryOne[userstructs.UserStruct](r.Context(), u.Dh, "SelectUserById", userId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get user or more then one user"))
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, user) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handles getting the user by tenant and user name
func (u *UserHandler) HandleGettingUserByTenantIdAndUserName(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "User | HandleGettingUserByTenantIdAndUserName")

	tenantId, err := httputils.IntPathParam(r, "tenant_id")
	userName := r.PathValue("user_name")

	if err != nil || userName == "" {
		fail(http.StatusBadRequest, errors.New("Failed to get user name or tenant id"))
		return
	}

	user, ok := database.QueryOne[userstructs.UserStruct](r.Context(), u.Dh, "SelectUserByTenantIdAndUserName", tenantId, userName)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting user or user is more then one"))
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, user) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handles the request for checking if a user has a tenant
func (u *UserHandler) HandleHasUserATenant(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "User | HandleHasUserATenant")

	userId, err := httputils.IntPathParam(r, "user_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	user, ok := database.QueryOne[userstructs.UserStruct](r.Context(), u.Dh, "SelectUserById", userId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get user or user is more then one"))
		return
	}

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
