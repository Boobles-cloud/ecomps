package handlers

import (
	"errors"
	"net/http"

	"ecomps.boobles.cloud/backend/database"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// ============= NOTE =============
// 	This needs the frontend auth
// ================================

// Handles getting all permissions for a user
func (u *UserHandler) HandleGettingUserPermissions(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Permission | HandleGettingUserPermission")

	userId, err := httputils.IntPathParam(r, "user_id")

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

	if !jsonutils.RespondWithJson(w, http.StatusOK, allPermissions) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handels getting permission by the given permission id
func (u *UserHandler) HandleGettingPermissionById(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Permission | HandleGettingPermissionById")

	permissionId, err := httputils.IntPathParam(r, "permission_id")

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

	if !jsonutils.RespondWithJson(w, http.StatusOK, wantedPermission) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handels getting all permissions by a language id
func (u *UserHandler) HandleGettingAllPermissionsByLanguageId(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Permission | HandleGettingAllPermissionsByLanguageId")

	languageId, err := httputils.IntPathParam(r, "language_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// Not to cause confusion:
	// We select all tenant actions here, because those are the real permissions.
	// A User gets access to a specific action he can do
	allPermissions, ok := database.QueryMany[userstructs.UserPermission](r.Context(), u.Dh, "SelectAllTenantActionsByLanguageId", []any{languageId})

	if !ok {
		fail(http.StatusInternalServerError, errors.New("Failed to get all permissions"))
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, allPermissions) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
