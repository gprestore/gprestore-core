package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gprestore/gprestore-core/internal/domain/user"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/pkg/handler"
	"github.com/gprestore/gprestore-core/pkg/variable"
	"github.com/markbates/goth/gothic"
)

type Handler struct {
	UserService *user.Service
}

func NewHandler(userService *user.Service) *Handler {
	return &Handler{
		UserService: userService,
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

	auth := &model.Auth{
		Action: variable.AUTH_ACTION_LOGIN,
		Token: &model.AuthToken{
			AccessToken:  gothUser.AccessToken,
			ExpiryAt:     &gothUser.ExpiresAt,
			RefreshToken: gothUser.RefreshToken,
		},
		Provider: gothUser.Provider,
	}

	filter := &model.UserFilter{
		Email: gothUser.Email,
	}

	user, err := h.UserService.FindOne(filter)
	if err != nil {
		auth.Action = variable.AUTH_ACTION_REGISTER

		input := &model.UserCreate{
			Username: "user" + gothUser.UserID,
			FullName: gothUser.Name,
			Email:    gothUser.Email,
			VerifyStatus: model.UserVerifyStatus{
				Email: true,
			},
			Image: gothUser.AvatarURL,
		}

		user, err = h.UserService.Create(input)
		if err != nil {
			handler.HandleError(w, err)
			return
		}
	}

	auth.User = user

	handler.SendSuccess(w, auth)
}
