package handler

import (
	"context"
	"errors"
	"net/http"
	"time"
	"uuid"

	"ecomps.boobles.cloud/backend/internal/middleware"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

func (t *TenantOAuthHandler) HandleOAuthCallback(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "[TenantOAuthHandler | HandleOAuthCallback] ")

	state := r.FormValue("state")
	stateCookie, err := r.Cookie(stateCookieKey)

	if err != nil || state != stateCookie.Value {
		fail(http.StatusBadRequest, errors.New("Invalid state"))
		return
	}

	// Get the provider and the appId from the path values
	// TODO: Dont forget to place them as callback url!
	oauthProvider := r.PathValue("provider")
	appId, err := httputils.IntPathParam(r, "app_id")

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   stateCookieKey,
		Value:  "",
		Path:   "/tenant/oauth/callback",
		MaxAge: -1,
	})

	code := r.FormValue("code")
	tenantId := ctx.Value(middleware.TenantIdContextKey).(uuid.UUID)

	if err := t.oauthService.AddToken(ctx, code, oauthProvider, tenantId, uint(appId)); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
