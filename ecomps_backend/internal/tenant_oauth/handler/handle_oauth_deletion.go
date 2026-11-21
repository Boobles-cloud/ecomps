package handler

import (
	"context"
	"net/http"
	"time"
	"uuid"

	"ecomps.boobles.cloud/backend/internal/middleware"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

func (t *TenantOAuthHandler) HandleTenantAppDeletion(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "[TenantOAuth | HandleTenantAppDeletion] ")

	tenantAppId, err := httputils.IntPathParam(r, "tenantapp_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	tenantId := ctx.Value(middleware.TenantIdContextKey).(uuid.UUID)

	if err := t.oauthService.DeleteOAuthAppForTenant(ctx, tenantId, uint(tenantAppId)); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
