package handlers

import (
	"context"
	"net/http"
	"time"

	"ecomps.boobles.cloud/backend/internal/auth/services"
	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handels the user logout
func (a *AuthHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Auth | HandleLogout")

	rqCookie, err := r.Cookie(services.AuthTokenCookieName)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if err := a.authService.DeleteAccessTokenByValue(ctx, rqCookie.Value); err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	// Send a new cookie so the old one gets deleted
	cookie := http.Cookie{
		Name:     services.AuthTokenCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
