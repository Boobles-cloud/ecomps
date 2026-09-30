package handlers

import (
	"context"
	"net/http"
	"time"

	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// ============= NOTE =============
// 	This needs the frontend auth
// ================================

// Handles getting all permissions for a user
func (ph *PermissionHandler) HandleGettingUserPermissions(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Permission | HandleGettingUserPermission")

	userId, err := httputils.IntPathParam(r, "user_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// Gets all permissions
	allPermissions, err := ph.permissionService.GetAllPermissionsForUserId(ctx, uint(userId))

	if !jsonutils.RespondWithJson(w, http.StatusOK, allPermissions) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handels getting permission by the given permission id
func (ph *PermissionHandler) HandleGettingPermissionById(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Permission | HandleGettingPermissionById")

	permissionId, err := httputils.IntPathParam(r, "permission_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// Not to cause confusion:
	// We select all tenant actions here, because those are the real permissions.
	// A User gets access to a specific action he can do
	wantedPermission, err := ph.permissionService.GetPermissionById(ctx, uint(permissionId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, wantedPermission) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handels getting all permissions by a language id
func (ph *PermissionHandler) HandleGettingAllPermissionsByLanguageId(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Permission | HandleGettingAllPermissionsByLanguageId")

	languageId, err := httputils.IntPathParam(r, "language_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	// Not to cause confusion:
	// We select all tenant actions here, because those are the real permissions.
	// A User gets access to a specific action he can do
	allPermissions, err := ph.permissionService.GetAllPermissionsByLanguageId(ctx, uint(languageId))

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	if !jsonutils.RespondWithJson(w, http.StatusOK, allPermissions) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
