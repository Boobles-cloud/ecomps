package middleware

import (
	"net/http"

	"boobles.cloud/backend/database"
	tenantstructs "boobles.cloud/backend/internal/tenant/tenant_structs"
)

// This middleware checks if a user has admin
func CheckAdminMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tenantId := ctx.Value(TenantIdContextKey).(uint)

		tenant, ok := database.QueryDatabase[tenantstructs.Tenant]("SelectTenantById", []any{tenantId})

		if !ok || len(tenant) != 1 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userId := ctx.Value(UserIdContextKey).(uint)

		if tenant[0].IsUserAdmin(userId) {
			h.ServeHTTP(w, r)
		}

		w.WriteHeader(http.StatusUnauthorized)
	})
}
