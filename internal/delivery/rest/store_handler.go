package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/service"
	"github.com/gprestore/gprestore-core/pkg/handler"
	"github.com/gprestore/gprestore-core/pkg/variable"
)

type StoreHandler struct {
	service *service.StoreService
}

func NewStoreHandler(service *service.StoreService) *StoreHandler {
	return &StoreHandler{
		service: service,
	}
}

func (h *StoreHandler) CreateStore(w http.ResponseWriter, r *http.Request) {
	var input *model.StoreCreate
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	authClaims, ok := r.Context().Value(variable.ContextKeyUser).(*model.AuthAccessTokenClaims)
	if !ok {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	// Secure Create Store
	// Author can create store for himself
	// Admin can create store for anyone
	if input.AuthorID != "" {
		if !(input.AuthorID == authClaims.UserId || authClaims.Role == variable.ROLE_ADMIN) {
			handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}
	} else {
		input.AuthorID = authClaims.UserId
	}

	store, err := h.service.Create(input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, store)
}

func (h *StoreHandler) UpdateStoreById(w http.ResponseWriter, r *http.Request) {
	var input *model.StoreUpdate
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	filter := &model.StoreFilter{
		Id: r.PathValue("id"),
	}

	store, err := h.service.FindOne(filter)
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
	if !(store.AuthorID == authClaims.UserId || authClaims.Role == variable.ROLE_ADMIN) {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	// Secure Change Author
	// Only Admin can change the author
	if input.AuthorID != "" && authClaims.Role != variable.ROLE_ADMIN {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	// Secure Change Badges
	// Only Admin can change the badges
	if input.Badges != nil && authClaims.Role != variable.ROLE_ADMIN {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	store, err = h.service.Update(filter, input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, store)
}

func (h *StoreHandler) DeleteStoreById(w http.ResponseWriter, r *http.Request) {
	filter := &model.StoreFilter{
		Id: r.PathValue("id"),
	}

	store, err := h.service.FindOne(filter)
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
	// Only Author or Admin can change the data
	if !(store.AuthorID == authClaims.UserId || authClaims.Role == variable.ROLE_ADMIN) {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	store, err = h.service.Delete(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, store)
}

func (h *StoreHandler) FindMany(w http.ResponseWriter, r *http.Request) {
	filter := &model.StoreFilter{
		Id:       r.URL.Query().Get("id"),
		Slug:     r.URL.Query().Get("slug"),
		AuthorID: r.URL.Query().Get("author_id"),
	}

	stores, err := h.service.FindMany(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, stores)
}

func (h *StoreHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	filter := &model.StoreFilter{
		Id:       r.URL.Query().Get("id"),
		Slug:     r.URL.Query().Get("slug"),
		AuthorID: r.URL.Query().Get("author_id"),
	}

	store, err := h.service.FindOne(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, store)
}
