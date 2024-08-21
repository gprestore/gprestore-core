package rest

import (
	"fmt"
	"net/http"

	"github.com/gprestore/gprestore-core/internal/service"
	"github.com/gprestore/gprestore-core/pkg/handler"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) OAuth(w http.ResponseWriter, r *http.Request) {
	provider := r.PathValue("provider")
	if provider == "" {
		handler.SendError(w, fmt.Errorf("provider is empty"), http.StatusBadRequest)
		return
	}

	q := r.URL.Query()
	q.Set("provider", provider)
	r.URL.RawQuery = q.Encode()

	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err == nil {
		handler.SendSuccess(w, gothUser)
		return
	}

	gothic.BeginAuthHandler(w, r)
}

func (h *AuthHandler) OAuthCallback(w http.ResponseWriter, r *http.Request) {
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	auth, err := h.service.LoginOrRegister(&gothUser)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "refreshToken",
		Value: auth.Token.RefreshToken,
		Path:  "/",
	})

	handler.SendSuccess(w, auth)
}

func (h *AuthHandler) CheckRefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := r.Cookie("refreshToken")
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	authToken, err := h.service.RefreshToken(refreshToken.Value)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, authToken)
}
