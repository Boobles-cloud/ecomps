package middleware

import (
	"net/http"
	"os"
	"strings"
)

// Used for frontend authentication
func FrontendAuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authToken := r.Header.Get("Authorization")

		if authToken == "" {
			w.WriteHeader(http.StatusBadRequest)
		}

		var tokenWithoutBaerer string

		// Replaces the bearer with empty string
		if strings.Contains(authToken, "bearer") {
			tokenWithoutBaerer = strings.ReplaceAll(authToken, "bearer ", "")
		} else {
			tokenWithoutBaerer = strings.ReplaceAll(authToken, "Bearer ", "")
		}

		apiKey := os.Getenv("API_Key")

		if tokenWithoutBaerer == apiKey {
			h.ServeHTTP(w, r)
		}

		w.WriteHeader(http.StatusUnauthorized)
	})
}
