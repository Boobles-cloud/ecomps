package middleware

import (
	"net/http"

	"boobles.cloud/backend/database"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
)

// This middleware checks if a user has admin
func CheckAdminMiddleware(dh *database.DbHandler) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			tenantId := ctx.Value(TenantIdContextKey).(uint)

			tenant, ok := database.QueryOne[tenantstructs.Tenant](ctx, dh, "SelectTenantById", []any{tenantId})

			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			userId := ctx.Value(UserIdContextKey).(uint)

			if tenant.IsUserAdmin(userId) {
				h.ServeHTTP(w, r)
			}

			w.WriteHeader(http.StatusUnauthorized)
		})
	}
}
