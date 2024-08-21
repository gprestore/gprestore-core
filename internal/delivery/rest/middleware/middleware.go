package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gprestore/gprestore-core/internal/model"
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

	return m.User(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(variable.ContextKeyUser).(*model.AuthAccessTokenClaims)
		if !ok {
			log.Println(user)
			handler.SendError(w, r, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}

		if user.Role != variable.ROLE_ADMIN {
			handler.SendError(w, r, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}))
}

func (m *Middleware) User(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		splitted := strings.Split(authorization, "Bearer ")
		if len(splitted) < 2 {
			handler.SendError(w, r, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}

		accessToken := splitted[1]
		refreshToken, err := r.Cookie("refreshToken")
		if err != nil {
			handler.SendError(w, r, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}

		if accessToken == "" {
			handler.SendError(w, r, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}

		accessTokenClaims, err := m.authService.ValidateAccessToken(accessToken)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				authToken, err := m.authService.RefreshToken(refreshToken.Value)
				if err != nil {
					handler.HandleError(w, r, err)
					return
				}
				accessToken = authToken.AccessToken
				accessTokenClaims, _ = m.authService.ValidateAccessToken(accessToken)
			} else {
				handler.HandleError(w, r, err)
				return
			}
		}

		ctxAccessToken := context.WithValue(context.Background(), variable.ContextKeyAccessToken, accessToken)
		ctxUser := context.WithValue(ctxAccessToken, variable.ContextKeyUser, accessTokenClaims)
		r = r.WithContext(ctxUser)

		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) Guest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
