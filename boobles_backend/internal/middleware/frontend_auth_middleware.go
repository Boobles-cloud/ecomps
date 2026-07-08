package middleware

import "net/http"

// Used for frontend authentication
func FrontendAuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
