package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/service"
	"github.com/gprestore/gprestore-core/pkg/handler"
	"github.com/gprestore/gprestore-core/pkg/variable"
)

type ItemHandler struct {
	service      *service.ItemService
	storeService *service.StoreService
}

func NewItemHandler(service *service.ItemService, storeService *service.StoreService) *ItemHandler {
	return &ItemHandler{
		service:      service,
		storeService: storeService,
	}
}

func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input *model.ItemCreate
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

	storeFilter := &model.StoreFilter{
		AuthorID: authClaims.UserId,
	}

	store, err := h.storeService.FindOne(storeFilter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	// Secure Create Item
	// Author can create item for himself
	// Admin can create item for anyone
	if input.StoreId != "" {
		if !(input.StoreId == store.Id.Hex() || authClaims.Role == variable.ROLE_ADMIN) {
			handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}
	} else {
		input.StoreId = store.Id.Hex()
	}

	item, err := h.service.Create(input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, item)
}

func (h *ItemHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	var input *model.ItemUpdate
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

	storeFilter := &model.StoreFilter{
		AuthorID: authClaims.UserId,
	}

	store, err := h.storeService.FindOne(storeFilter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	// Secure Change Data
	// Only Author or Admin can change the data
	if !(store.AuthorID == authClaims.UserId || authClaims.Role == variable.ROLE_ADMIN) {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	// Secure Change Store
	// Only Admin can change the store
	if input.StoreId != "" && authClaims.Role != variable.ROLE_ADMIN {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	filter := &model.ItemFilter{
		Id: r.PathValue("id"),
	}

	item, err := h.service.Update(filter, input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, item)
}

func (h *ItemHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	authClaims, ok := r.Context().Value(variable.ContextKeyUser).(*model.AuthAccessTokenClaims)
	if !ok {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	storeFilter := &model.StoreFilter{
		AuthorID: authClaims.UserId,
	}

	store, err := h.storeService.FindOne(storeFilter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	// Secure Delete Data
	// Only Author or Admin can delete the data
	if !(store.AuthorID == authClaims.UserId || authClaims.Role == variable.ROLE_ADMIN) {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	filter := &model.ItemFilter{
		Id: r.PathValue("id"),
	}

	item, err := h.service.Delete(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, item)
}

func (h *ItemHandler) FindMany(w http.ResponseWriter, r *http.Request) {
	filter := &model.ItemFilter{
		Id:      r.URL.Query().Get("id"),
		Slug:    r.URL.Query().Get("slug"),
		StoreId: r.URL.Query().Get("store_id"),
	}

	// Secure Find Items
	// Store ID Required
	if filter.StoreId == "" {
		handler.HandleError(w, variable.ErrItemFilterStoreId)
		return
	}

	items, err := h.service.FindMany(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, items)
}

func (h *ItemHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	filter := &model.ItemFilter{
		Id:      r.URL.Query().Get("id"),
		Slug:    r.URL.Query().Get("slug"),
		StoreId: r.URL.Query().Get("store_id"),
	}

	item, err := h.service.FindOne(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, item)
}
