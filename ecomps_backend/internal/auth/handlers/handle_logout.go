package handlers

import (
	"errors"
	"net/http"
	"time"

	httputils "ecomps.boobles.cloud/backend/utils/http_utils"
)

// Handels the user logout
func (ha *AuthHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {

	fail := httputils.NewFailHandler(w, "Auth | HandleLogout")

	rqCookie, err := r.Cookie(httputils.AuthTokenCookieName)

	if err != nil {
		fail(http.StatusBadRequest, err)
		return
	}

	if result := ha.Dh.ExecuteSQLStatement("DeleteUserAccestokenByValue", []any{rqCookie.Value}); !result.Ok {
		fail(http.StatusInternalServerError, errors.New("Failed deleting authtoken"))
		return
	}

	// Send a new cookie so the old one gets deleted
	cookie := http.Cookie{
		Name:     httputils.AuthTokenCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
