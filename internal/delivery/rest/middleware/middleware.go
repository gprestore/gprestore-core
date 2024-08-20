package middleware

import (
	"net/http"
	"strings"

	"github.com/gprestore/gprestore-core/internal/service"
	"github.com/gprestore/gprestore-core/pkg/handler"
	"github.com/gprestore/gprestore-core/pkg/variable"
)

type Middleware struct {
	authService *service.AuthService
}

func NewMiddleware(authService *service.AuthService) *Middleware {
	return &Middleware{
		authService: authService,
	}
}

func (m *Middleware) Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		splitted := strings.Split(authorization, "Bearer ")
		if len(splitted) < 2 {
			handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}

		accessToken := splitted[1]
		if accessToken == "" {
			handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}

		user, err := m.authService.ValidateAccessToken(accessToken)
		if err != nil {
			handler.HandleError(w, err)
			return
		}

		if user.Role != variable.ROLE_ADMIN {
			handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) User(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		splitted := strings.Split(authorization, "Bearer ")
		if len(splitted) < 2 {
			handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}

		accessToken := splitted[1]
		if accessToken == "" {
			handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}

		_, err := m.authService.ValidateAccessToken(accessToken)
		if err != nil {
			handler.HandleError(w, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) Guest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
