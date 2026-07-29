package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

// Creates the middleware stack
func CreateNewMiddlewareStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := range xs {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}
