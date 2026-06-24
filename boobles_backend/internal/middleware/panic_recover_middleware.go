package middleware

import (
	"net/http"

	"boobles.cloud/backend/logging"
)

func PanicRecoverMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				logging.Log(logging.Error, err.(string))
			}
		}()

		h.ServeHTTP(w, r)
	})
}
