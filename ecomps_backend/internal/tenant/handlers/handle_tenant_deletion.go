package handlers

import (
	"context"
	"net/http"
	"time"

	"ecomps.boobles.cloud/backend/internal/middleware"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handels the deletion of a tenant
func (t *TenantHandler) HandleTenantDeletion(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Tenant | HandleTenantDeletion")

	userId := ctx.Value(middleware.UserIdContextKey).(int)
	tenantId := ctx.Value(middleware.TenantIdContextKey).(int)

	if err := t.tenantService.Delete(ctx, uint(userId), uint(tenantId)); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
