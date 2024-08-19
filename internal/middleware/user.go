package middleware

import (
	"log"
	"net/http"
)

func User(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware User")

		next.ServeHTTP(w, r)
	})
}
