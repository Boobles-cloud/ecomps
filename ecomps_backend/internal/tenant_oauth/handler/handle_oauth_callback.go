package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

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

	oauthProvider := r.PathValue("provider")

	if oauthProvider == "" {
		fail(http.StatusBadRequest, errors.New("OAuthprovider is empty"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   stateCookieKey,
		Value:  "",
		Path:   "/tenant/oauth/callback",
		MaxAge: -1,
	})

	code := r.FormValue("code")

	if code == "" {
		fail(http.StatusBadRequest, errors.New("Code not found"))
		return
	}

	if err := t.oauthService.AddToken(ctx, code, oauthProvider); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
