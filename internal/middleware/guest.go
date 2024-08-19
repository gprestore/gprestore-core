package middleware

import (
	"log"
	"net/http"
)

func Guest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware Guest")

		next.ServeHTTP(w, r)
	})
}
