package middleware

import (
	"net/http"

	"ecomps.boobles.cloud/backend/logging"
)

// Just for logging everything
func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logMsg := "Request from: " + r.RemoteAddr + " ; Target: " + r.RequestURI + " ; Method: " + r.Method
		logging.Log(logging.Information, logMsg)
		h.ServeHTTP(w, r)
	})
}
