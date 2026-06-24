package middleware

import "net/http"

// Checks if the user has a permission
// NOTE: Only use this middleware after checking if a user is authenticated!!!!
func PermissionMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement meeeee
	})
}
