package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gprestore/gprestore-core/pkg/handler"
	"github.com/markbates/goth/gothic"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) OAuth(w http.ResponseWriter, r *http.Request) {
	log.Println("A")
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

func (h *Handler) OAuthCallback(w http.ResponseWriter, r *http.Request) {
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

	handler.SendSuccess(w, auth)
}
