package handlers

import (
	"net/http"
	"strings"
	"time"

	"boobles.cloud/backend/database"
	"boobles.cloud/backend/logging"
)

// Handels the user logout
func HandleLogout(w http.ResponseWriter, r *http.Request) {

	fail := func(status int, err error) {

		if err != nil {
			logging.Log(logging.Error, "[Auth | HandleLogout]"+err.Error())
		}

		w.WriteHeader(status)
	}

	token := r.Header.Get("Authorization")

	if token == "" {
		fail(http.StatusBadRequest, nil)
	}

	var tokenWithoutBaerer string
	if strings.Contains(token, "bearer") {
		tokenWithoutBaerer = strings.ReplaceAll(token, "bearer ", "")
	} else {
		tokenWithoutBaerer = strings.ReplaceAll(token, "Bearer ", "")
	}

	if result := database.ExecuteSQLStatement("DeleteUserAccestokenByValue", database.Delete, []any{tokenWithoutBaerer}); !result.Ok {
		fail(http.StatusInternalServerError, nil)
	}

	// Send a new cookie so the old one gets deleted
	cookie := http.Cookie{
		Name:     "AuthTokenBoobles",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
