package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	permissionstructs "ecomps.boobles.cloud/backend/internal/permission/permission_structs"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
)

// Handels the adding of a new permission for the user
// NOTE: for this to work, use the [CheckAdminMiddleware]
// Only admins can access it
func (ph *PermissionHandler) HandleAddingNewUserPermission(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "User | HandleAddingNewUserPermission")

	userPermission, err := jsonutils.JsonDeserilizeHttpRequestBody[permissionstructs.Permission](r)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if err := ph.permissionService.AsignUserPermssion(ctx, userPermission.UserId, userPermission.PermissionId); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Handels the removing of a user permission
// NOTE: for this to work, use the [CheckAdminMiddleware]
// Only admins can access it.
func (ph *PermissionHandler) HandleRemovingUserPermission(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "User | HandleRemovinUserPermission")

	permissionId, err := httputils.IntPathParam(r, "permission_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
	}

	userId, err := httputils.IntPathParam(r, "user_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if err := ph.permissionService.RemoveUserPermission(ctx, uint(userId), uint(permissionId)); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	fail(http.StatusBadRequest, errors.New("NOT IMPLEMENTED"))
}
