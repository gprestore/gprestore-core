package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gprestore/gprestore-core/internal/config"
	"github.com/gprestore/gprestore-core/internal/domain/user"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/pkg/handler"
	"github.com/gprestore/gprestore-core/pkg/variable"
	"github.com/spf13/viper"
)

type Handler struct {
	UserService *user.Service
}

func NewHandler(userService *user.Service) *Handler {
	return &Handler{
		UserService: userService,
	}
}

func (h *Handler) LoginGoogle(w http.ResponseWriter, r *http.Request) {
	url := config.AppConfig.OAuthGoogle.AuthCodeURL(viper.GetString("oauth.state"))
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	if state != viper.GetString("oauth.state") {
		handler.SendError(w, fmt.Errorf("unauthorized"), http.StatusUnauthorized)
		return
	}

	code := r.URL.Query().Get("code")
	token, err := config.AppConfig.OAuthGoogle.Exchange(context.Background(), code)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	var authGoogle *model.AuthGoogle
	err = json.NewDecoder(resp.Body).Decode(&authGoogle)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	filter := &model.UserFilter{
		Email: authGoogle.Email,
	}

	auth := &model.Auth{
		Action: variable.AUTH_ACTION_LOGIN,
		Token:  token,
	}

	user, err := h.UserService.FindOne(filter)
	if err != nil {
		auth.Action = variable.AUTH_ACTION_REGISTER

		input := &model.UserCreate{
			Username: "user" + authGoogle.Id,
			FullName: authGoogle.Name,
			Email:    authGoogle.Email,
			VerifyStatus: model.UserVerifyStatus{
				Email: authGoogle.VerifiedEmail,
			},
			Image: authGoogle.Picture,
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
