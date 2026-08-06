package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"boobles.cloud/backend/database"
	"boobles.cloud/backend/internal/middleware"
	userstructs "boobles.cloud/backend/internal/user/user_structs"
	"boobles.cloud/backend/logging"
)

// ============= NOTE =============
// 	This needs the frontend auth
// ================================

// Handles the getting the user by the accesstoken val
func (u *UserHandler) HandleGettingUserByAuthTokenVal(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(status)
	}

	authToken := r.PathValue("authtoken")

	if authToken == "" {
		fail(http.StatusBadRequest, errors.New("Failed getting authtoken or token is null"))
	}

	user, ok := database.QueryOne[userstructs.UserStruct](r.Context(), u.Dh, "SelectUserByAuthToken", []any{authToken})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting user or user is more then one"))
	}

	jsonData, err := json.Marshal(user)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

// Handles getting the user by the user id
func (u *UserHandler) HandleGettingUserById(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(status)
	}

	userId := r.Context().Value(middleware.UserIdContextKey).(int)

	user, ok := database.QueryOne[userstructs.UserStruct](r.Context(), u.Dh, "SelectUserById", []any{userId})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get user or more then one user"))
	}

	jsonData, err := json.Marshal(user)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

// Handles getting the user by tenant and user name
func (u *UserHandler) HandleGettingUserByTenantIdAndUserName(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(status)
	}

	tenantIdString := r.PathValue("tenant-id")
	userName := r.PathValue("user-name")

	if tenantIdString == "" || userName == "" {
		fail(http.StatusBadRequest, errors.New("Failed to get user name or tenant id"))
	}

	tenantId, err := strconv.Atoi(tenantIdString)

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	user, ok := database.QueryOne[userstructs.UserStruct](r.Context(), u.Dh, "SelectUserByTenantIdAndUserName", []any{tenantId, userName})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting user or user is more then one"))
	}

	jsonData, err := json.Marshal(user)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)

}

// Handles the request for checking if a user has a tenant
func (u *UserHandler) HandleHasUserATenant(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, err.Error())
		w.WriteHeader(status)
	}

	userId := r.Context().Value(middleware.UserIdContextKey).(int)

	user, ok := database.QueryOne[userstructs.UserStruct](r.Context(), u.Dh, "SelectUserById", []any{userId})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get user or user is more then one"))
	}

	if user.UserHasTenant && user.TenantId != 0 {
		jsonData, err := json.Marshal("Tenant = true")

		if err != nil {
			fail(http.StatusInternalServerError, err)
		}

		w.Write(jsonData)
		w.WriteHeader(http.StatusOK)
	}

	jsonData, err := json.Marshal("Tenant = false")

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)

}
