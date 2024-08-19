package middleware

import (
	"log"
	"net/http"
)

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware Admin")

		next.ServeHTTP(w, r)
	})
}
