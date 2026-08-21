package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// Creates the middleware stack
func CreateNewMiddlewareStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			next = xs[i](next)
		}
		return next
	}
}
