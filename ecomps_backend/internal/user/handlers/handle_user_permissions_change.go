package handlers

import (
	"errors"
	"net/http"

	"ecomps.boobles.cloud/backend/database"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels the adding of a new permission for the user
// NOTE: for this to work, use the [CheckAdminMiddleware]
// Only admins can access it
func (u *UserHandler) HandleAddingNewUserPermission(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "User | HandleAddingNewUserPermission")

	userPermission, err := jsonutils.JsonDeserilizeHttpRequestBody[userstructs.UserPermission](r)

	if err != nil {
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
		if allPermissions[i].PermissionName == userPermission.PermissionName {
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

	fail := httputils.NewFailHandler(w, "User | HandleRemovinUserPermission")

	fail(http.StatusBadRequest, errors.New("NOT IMPLEMENTED"))
}
