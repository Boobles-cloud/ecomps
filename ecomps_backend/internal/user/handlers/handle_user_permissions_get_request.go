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

// Handles getting all permissions for a user
func (u *UserHandler) HandleGettingUserPermissions(w http.ResponseWriter, r *http.Request) {
	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Permission | HandleGettingUserPemissions] "+err.Error())
		}

		w.WriteHeader(status)
	}

	userId, err := strconv.Atoi(r.PathValue("user_id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// Gets the wanted user from the database
	wantedUser, ok := database.QueryOne[userstructs.UserStruct](r.Context(), u.Dh, "SelectUserById", userId)

	// Checks for err and if its only one
	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get user from database"))
		return
	}

	// Gets all permissions
	allPermissions, ok := wantedUser.GetPermissionsByUser(r.Context(), u.Dh)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get user permissions"))
		return
	}

	jsonData, err := json.Marshal(allPermissions)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Handels getting permission by the given permission id
func (u *UserHandler) HandleGettingPermissionById(w http.ResponseWriter, r *http.Request) {
	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Permission | HandleGettingPermissionById] "+err.Error())
		}

		w.WriteHeader(status)
	}

	permissionId, err := strconv.Atoi(r.PathValue("permission_id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// Not to cause confusion:
	// We select all tenant actions here, because those are the real permissions.
	// A User gets access to a specific action he can do
	wantedPermission, ok := database.QueryOne[userstructs.UserPermission](r.Context(), u.Dh, "SelectTenantActionById", permissionId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get permission from database"))
		return
	}

	jsonData, err := json.Marshal(wantedPermission)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Handels getting all permissions
func (u *UserHandler) HandleGettingAllPermissions(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Permission | HandleGettingAllPermissions] "+err.Error())
		}

		w.WriteHeader(status)
	}

	// Not to cause confusion:
	// We select all tenant actions here, because those are the real permissions.
	// A User gets access to a specific action he can do
	allPermissions, ok := database.QueryMany[userstructs.UserPermission](r.Context(), u.Dh, "SelectAllTenantActions", []any{})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get all permissions"))
		return
	}

	jsonData, err := json.Marshal(allPermissions)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
