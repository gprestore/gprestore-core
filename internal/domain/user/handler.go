package user

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/pkg/handler"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	var input *model.UserUpdate
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	filter := &model.UserFilter{
		Id: r.PathValue("id"),
	}

	user, err := h.service.Update(filter, input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, user)
}

func (h *Handler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	filter := &model.UserFilter{
		Id: r.PathValue("id"),
	}

	user, err := h.service.Delete(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, user)
}

func (h *Handler) FindMany(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) FindOne(w http.ResponseWriter, r *http.Request) {
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
