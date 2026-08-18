package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"ecomps.boobles.cloud/backend/database"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	"ecomps.boobles.cloud/backend/utils/logging"
)

// Handels the adding of a new permission for the user
// NOTE: for this to work, use the [CheckAdminMiddleware]
// Only admins can access it
func (u *UserHandler) HandleAddingNewUserPermission(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {
		logging.Log(logging.Error, "[Permission | HandleAddingNewUserPermission] "+err.Error())
		w.WriteHeader(status)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	userPermission := userstructs.UserPermission{}

	if err := json.Unmarshal(body, &userPermission); err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	user, ok := database.QueryOne[userstructs.UserStruct](r.Context(), u.Dh, "SelectUserById", userPermission.UserId)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting user"))
		return
	}

	allPermissions, ok := user.GetPermissionsByUser(r.Context(), u.Dh)

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed getting user permissions"))
		return
	}

	for i := range allPermissions {
		if allPermissions[i].PermissionName == user.UserName {
			w.WriteHeader(http.StatusOK)
		}
	}

	if _, ok := userPermission.SetNewPermission(u.Dh); !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to create user permission"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Handels the removing of a user permission
// NOTE: for this to work, use the [CheckAdminMiddleware]
// Only admins can access it.
func (u *UserHandler) HandleRemovingUserPermission(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Permission | HandleRemovingNewUserPermission] "+err.Error())
		}

		w.WriteHeader(status)
	}

	fail(http.StatusBadRequest, errors.New("NOT IMPLEMENTED"))
	return
}
