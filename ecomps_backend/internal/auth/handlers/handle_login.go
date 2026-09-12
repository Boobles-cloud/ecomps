package handlers

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"

	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Creates a new JWT for the given user.
// The pw from the user, is encrypted via the frontend.
func (a *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	fail := httputils.NewFailHandler(w, "Auth | HandleLogin")

	basicAuth := r.Header.Get("Authorization")

	if basicAuth == "" {
		fail(http.StatusBadRequest, errors.New("Failed getting Authorization header"))
		return
	}

	authPWUser := strings.ReplaceAll(basicAuth, "Basic ", "")

	encodedAuthPWUser, err := base64.StdEncoding.DecodeString(authPWUser)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	authSplitet := strings.Split(string(encodedAuthPWUser), ":")

	if len(authSplitet) != 2 {
		fail(http.StatusBadRequest, errors.New("Failed spliting basic pw"))
		return
	}

	user, err := a.userService.GetUserByUserNameAndPw(ctx, authSplitet[0], authSplitet[1])

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	cookie, err := a.authService.CreateAuthCookie(ctx, user.UserId, user.TenantId)

	if err != nil {
		fail(http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
