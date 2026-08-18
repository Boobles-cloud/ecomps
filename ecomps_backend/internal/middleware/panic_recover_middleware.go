package middleware

import (
	"fmt"
	"net/http"

	"ecomps.boobles.cloud/backend/logging"
)

// Used to log any panics our program does
func PanicRecoverMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				logging.Log(logging.Error, fmt.Sprintf("%v", err))
			}
		}()

		h.ServeHTTP(w, r)
	})
}
