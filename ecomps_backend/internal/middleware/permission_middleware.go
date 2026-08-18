package middleware

import (
	"errors"
	"net/http"

	"ecomps.boobles.cloud/backend/database"
	tenantstructs "ecomps.boobles.cloud/backend/internal/tenant/tenant_structs"
	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
	"ecomps.boobles.cloud/backend/logging"
)

const (
	HeaderPermissionVal = "ActionName"
)

// Checks if the user has a permission
// NOTE: Only use this middleware after checking if a user is authenticated!!!!
func PermissionMiddleware(dh *database.DbHandler) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			fail := func(code int, err error) {
				logging.Log(logging.Error, "[middlerware | PermissionMiddleware] "+err.Error())
				w.WriteHeader(code)
				w.Write([]byte(err.Error()))
			}

			// TODO: maybe in the future implement view permission stuff ;)
			if r.Method != "POST" {
				h.ServeHTTP(w, r)
			}

			userId, okUser := r.Context().Value(UserIdContextKey).(int)
			tenantId, okTenant := r.Context().Value(TenantIdContextKey).(int)

			if !okUser || !okTenant {
				fail(http.StatusBadRequest, errors.New("Failed getting user or tenant id from context"))
				return
			}

			tenant, ok := database.QueryOne[tenantstructs.Tenant](r.Context(), dh, "SelectTenantById", tenantId)

			if !ok {
				fail(http.StatusBadRequest, errors.New("User does not have a tenant"))
				return
			}

			if tenant.IsUserAdmin(uint(userId)) {
				h.ServeHTTP(w, r)
			}

			permissionName := r.Header.Get(HeaderPermissionVal)

			if permissionName == "" {
				fail(http.StatusBadRequest, errors.New("No action given"))
				return
			}

			permission, ok := database.QueryOne[userstructs.UserPermission](r.Context(), dh, "SelectPermissionByName", permissionName)

			// If the permission isn´t ok we return an unauthorized
			if !ok || permission.PermissionName != permissionName {
				fail(http.StatusUnauthorized, errors.New("Failed getting one or more then one permission"))
				return
			}

			// Else serve it
			h.ServeHTTP(w, r)
		})
	}
}
