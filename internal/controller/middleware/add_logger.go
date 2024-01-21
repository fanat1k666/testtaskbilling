package middleware

import (
	"TestTask/pkg/log"
	"net/http"
)

func AddLogger(next http.HandlerFunc, l log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(log.AddToContext(r.Context(), l))
		next.ServeHTTP(w, r)
	}
}
