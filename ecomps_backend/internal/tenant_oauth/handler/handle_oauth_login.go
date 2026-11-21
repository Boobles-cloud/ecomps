package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

func (t *TenantOAuthHandler) HandleOAuthLoginRequest(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	provider := r.PathValue("provider_key")

	fail := httputils.NewFailHandler(w, "[TenantOAuht | HandleOAuthLoginRequest")

	state := generateStateCookie(w)

	url, err := t.oauthService.GetAuthUrlByProviderKey(ctx, provider, state)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func generateStateCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(10 * time.Minute)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: stateCookieKey, Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)
	return state
}
