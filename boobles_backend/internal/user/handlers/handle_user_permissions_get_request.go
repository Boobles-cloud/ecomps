package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"boobles.cloud/backend/database"
	userstructs "boobles.cloud/backend/internal/user/user_structs"
	"boobles.cloud/backend/logging"
)

// Handles getting all permissions for a user
func HandleGettingUserPermissions(w http.ResponseWriter, r *http.Request) {
	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Permission | HandleGettingUserPemissions] "+err.Error())
		}

		w.WriteHeader(status)
	}

	userId, err := strconv.Atoi(r.URL.Query().Get("user-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	// Gets the wanted user from the database
	wantedUser, ok := database.QueryDatabase[userstructs.UserStruct]("SelectUserById", []any{userId})

	// Checks for err and if its only one
	if !ok && len(wantedUser) != 1 {
		fail(http.StatusInternalServerError, errors.New("Failed to get user from database"))
	}

	// Gets all permissions
	allPermissions, ok := wantedUser[0].GetPermissionsByUser()

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get user permissions"))
	}

	jsonData, err := json.Marshal(allPermissions)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

// Handels getting permission by the given permission id
func HandleGettingPermissionById(w http.ResponseWriter, r *http.Request) {
	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Permission | HandleGettingPermissionById] "+err.Error())
		}

		w.WriteHeader(status)
	}

	permissionId, err := strconv.Atoi(r.URL.Query().Get("permission-id"))

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	// Not to cause confusion:
	// We select all tenant actions here, because those are the real permissions.
	// A User gets access to a specific action he can do
	wantedPermission, ok := database.QueryDatabase[userstructs.UserPermission]("SelectTenantActionById", []any{permissionId})

	if !ok && len(wantedPermission) != 1 {
		fail(http.StatusInternalServerError, errors.New("Failed to get permission from database"))
	}

	jsonData, err := json.Marshal(wantedPermission)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

// Handels getting all permissions
func HandleGettingAllPermissions(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Permission | HandleGettingAllPermissions] "+err.Error())
		}

		w.WriteHeader(status)
	}

	// Not to cause confusion:
	// We select all tenant actions here, because those are the real permissions.
	// A User gets access to a specific action he can do
	allPermissions, ok := database.QueryDatabase[userstructs.UserPermission]("SelectAllTenantActions", []any{})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get all permissions"))
	}

	jsonData, err := json.Marshal(allPermissions)

	if err != nil {
		fail(http.StatusInternalServerError, err)
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}
