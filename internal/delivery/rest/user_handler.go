package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/service"
	"github.com/gprestore/gprestore-core/pkg/handler"
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
		handler.HandleError(w, r, err)
		return
	}

	user, err := h.service.Create(input)
	if err != nil {
		handler.HandleError(w, r, err)
		return
	}

	handler.SendSuccess(w, r, user)
}

func (h *UserHandler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	var input *model.UserUpdate
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		handler.HandleError(w, r, err)
		return
	}

	filter := &model.UserFilter{
		Id: r.PathValue("id"),
	}

	user, err := h.service.Update(filter, input)
	if err != nil {
		handler.HandleError(w, r, err)
		return
	}

	handler.SendSuccess(w, r, user)
}

func (h *UserHandler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	filter := &model.UserFilter{
		Id: r.PathValue("id"),
	}

	user, err := h.service.Delete(filter)
	if err != nil {
		handler.HandleError(w, r, err)
		return
	}

	handler.SendSuccess(w, r, user)
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
		handler.HandleError(w, r, err)
		return
	}

	handler.SendSuccess(w, r, users)
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
		handler.HandleError(w, r, err)
		return
	}

	handler.SendSuccess(w, r, user)
}
