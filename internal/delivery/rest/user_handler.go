package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/service"
	"github.com/gprestore/gprestore-core/pkg/handler"
	"github.com/gprestore/gprestore-core/pkg/variable"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input *model.UserCreate
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	user, err := h.service.Create(input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, user)
}

func (h *UserHandler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	var input *model.UserUpdate
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	filter := &model.UserFilter{
		Id: r.PathValue("id"),
	}

	user, err := h.service.FindOne(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	authClaims, ok := r.Context().Value(variable.ContextKeyUser).(*model.AuthAccessTokenClaims)
	if !ok {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	// Secure Change Data
	// Only Author or Admin can change the data
	if !(user.Id.Hex() == authClaims.UserId || authClaims.Role == variable.ROLE_ADMIN) {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	user, err = h.service.Update(filter, input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, user)
}

func (h *UserHandler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	filter := &model.UserFilter{
		Id: r.PathValue("id"),
	}

	user, err := h.service.FindOne(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	authClaims, ok := r.Context().Value(variable.ContextKeyUser).(*model.AuthAccessTokenClaims)
	if !ok {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	// Secure Delete Data
	// Only Author or Admin can delete the data
	if !(user.Id.Hex() == authClaims.UserId || authClaims.Role == variable.ROLE_ADMIN) {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	user, err = h.service.Delete(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, user)
}

func (h *UserHandler) FindMany(w http.ResponseWriter, r *http.Request) {
	filter := &model.UserFilter{
		Id:       r.URL.Query().Get("id"),
		Username: r.URL.Query().Get("username"),
		Email:    r.URL.Query().Get("email"),
		Phone:    r.URL.Query().Get("phone"),
	}
	if filter.Phone != "" {
		filter.Phone = "+" + strings.Trim(filter.Phone, " ")
	}

	users, err := h.service.FindMany(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, users)
}

func (h *UserHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	filter := &model.UserFilter{
		Id:       r.URL.Query().Get("id"),
		Username: r.URL.Query().Get("username"),
		Email:    r.URL.Query().Get("email"),
		Phone:    r.URL.Query().Get("phone"),
	}
	if filter.Phone != "" {
		filter.Phone = "+" + strings.Trim(filter.Phone, " ")
	}

	user, err := h.service.FindOne(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, user)
}
