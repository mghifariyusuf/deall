package middleware

import (
	"deall/cmd/lib/customError"
	"deall/pkg/response"
	"net/http"
)

func RequiresCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Request-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")

		if r.Method == "OPTIONS" {
			response.Error(w, customError.ErrNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
