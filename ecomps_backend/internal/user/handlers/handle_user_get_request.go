package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"ecomps.boobles.cloud/backend/database"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	"ecomps.boobles.cloud/backend/logging"
)

// ============= NOTE =============
// 	This needs the frontend auth
// ================================

// Handles the getting the user by the accesstoken val
func (u *UserHandler) HandleGettingUserByAuthTokenVal(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[User | HanldeGettingUserByAuthTokenVal]"+err.Error())
		w.WriteHeader(status)
	}

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

	jsonData, err := json.Marshal(user)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Handles getting the user by the user id
func (u *UserHandler) HandleGettingUserById(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[User | HandleGettingUserById] "+err.Error())
		w.WriteHeader(status)
	}

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

	jsonData, err := json.Marshal(user)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Handles getting the user by tenant and user name
func (u *UserHandler) HandleGettingUserByTenantIdAndUserName(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[User | HandleGettingUserByTenantIdAndUserName] "+err.Error())
		w.WriteHeader(status)
	}

	tenantIdString := r.PathValue("tenant-id")
	userName := r.PathValue("user_name")

	if tenantIdString == "" || userName == "" {
		fail(http.StatusBadRequest, errors.New("Failed to get user name or tenant id"))
		return
	}

	tenantId, err := strconv.Atoi(tenantIdString)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	user, ok := database.QueryOne[userstructs.UserStruct](r.Context(), u.Dh, "SelectUserByTenantIdAndUserName", tenantId, userName)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting user or user is more then one"))
		return
	}

	jsonData, err := json.Marshal(user)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Handles the request for checking if a user has a tenant
func (u *UserHandler) HandleHasUserATenant(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[User | HandleHasUserATenant] "+err.Error())
		w.WriteHeader(status)
	}

	userId, err := strconv.Atoi(r.PathValue("user_id"))

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
		jsonData, err := json.Marshal("Tenant = true")

		if err != nil {
			fail(http.StatusInternalServerError, err)
			return
		}

		w.Write(jsonData)
		w.WriteHeader(http.StatusOK)
	}

	jsonData, err := json.Marshal("Tenant = false")

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
