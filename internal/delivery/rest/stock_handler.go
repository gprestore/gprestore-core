package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/service"
	"github.com/gprestore/gprestore-core/pkg/handler"
	"github.com/gprestore/gprestore-core/pkg/variable"
)

type StockHandler struct {
	service      *service.StockService
	storeService *service.StoreService
}

func NewStockHandler(service *service.StockService, storeService *service.StoreService) *StockHandler {
	return &StockHandler{
		service:      service,
		storeService: storeService,
	}
}

func (h *StockHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input *model.StockCreate
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

	// Secure Create Stock
	// Author can create stock for himself
	// Admin can create stock for anyone
	if input.StoreId != "" {
		if !(input.StoreId == store.Id.Hex() || authClaims.Role == variable.ROLE_ADMIN) {
			handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}
	} else {
		input.StoreId = store.Id.Hex()
	}

	stock, err := h.service.Create(input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, stock)
}

func (h *StockHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	var input *model.StockUpdate
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

	filter := &model.StockFilter{
		Id: r.PathValue("id"),
	}

	stock, err := h.service.Update(filter, input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, stock)
}

func (h *StockHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	filter := &model.StockFilter{
		Id:      r.URL.Query().Get("id"),
		StoreId: r.URL.Query().Get("store_id"),
		ItemId:  r.URL.Query().Get("item_id"),
	}

	stock, err := h.service.FindOne(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, stock)
}
