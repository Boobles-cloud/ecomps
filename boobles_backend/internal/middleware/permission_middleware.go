package middleware

import (
	"errors"
	"net/http"

	"boobles.cloud/backend/database"
	userstructs "boobles.cloud/backend/internal/user/user_structs"
	"boobles.cloud/backend/logging"
)

const (
	HeaderPermissionVal = "ActionName"
)

// Checks if the user has a permission
// NOTE: Only use this middleware after checking if a user is authenticated!!!!
func PermissionMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fail := func(code int, err error) {
			logging.Log(logging.Error, "[middlerware | PermissionMiddleware] "+err.Error())
			w.WriteHeader(code)
		}

		// TODO: implement getting all permissions from cache

		// TODO: maybe in the future implement view permission stuff ;)
		if r.Method != "POST" {
			h.ServeHTTP(w, r)
		}

		permissionName := r.Header.Get(HeaderPermissionVal)

		if permissionName == "" {
			fail(http.StatusBadRequest, errors.New("No action given"))
		}

		permission, ok := database.QueryDatabase[userstructs.UserPermission]("SelectPermissionByName", []any{permissionName})

		// If the permission is more then one or it isn´t ok we return an unauthorized
		if !ok || len(permission) != 1 {
			fail(http.StatusUnauthorized, errors.New("Failed getting or more then one permission"))
		}

		// Else serve it
		h.ServeHTTP(w, r)
	})
}
