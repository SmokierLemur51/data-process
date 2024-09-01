package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(xm ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xm) - 1; i >= 0; i-- {
			x := xm[i]
			next = x(next)
		}

		return next
	}
}
