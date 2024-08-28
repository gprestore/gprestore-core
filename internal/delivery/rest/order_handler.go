package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/service"
	"github.com/gprestore/gprestore-core/pkg/handler"
	"github.com/gprestore/gprestore-core/pkg/variable"
)

type OrderHandler struct {
	service      *service.OrderService
	storeService *service.StoreService
}

func NewOrderHandler(service *service.OrderService, storeService *service.StoreService) *OrderHandler {
	return &OrderHandler{
		service:      service,
		storeService: storeService,
	}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input *model.OrderCreate
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	authClaims, ok := r.Context().Value(variable.ContextKeyUser).(*model.AuthAccessTokenClaims)
	if ok {
		storeFilter := &model.StoreFilter{
			AuthorID: authClaims.UserId,
		}

		store, err := h.storeService.FindOne(storeFilter)
		if err != nil {
			handler.HandleError(w, err)
			return
		}

		// Secure Create Order
		// Author can't create order from his own store
		if input.StoreId == store.Id.Hex() {
			handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
			return
		}
	}

	order, err := h.service.Create(input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, order)
}

func (h *OrderHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	var input *model.OrderUpdate
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

	// Secure Change Data
	// Only Admin can change the data
	if authClaims.Role != variable.ROLE_ADMIN {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	filter := &model.OrderFilter{
		Id: r.PathValue("id"),
	}

	order, err := h.service.Update(filter, input)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, order)
}

func (h *OrderHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	authClaims, ok := r.Context().Value(variable.ContextKeyUser).(*model.AuthAccessTokenClaims)
	if !ok {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	// Secure Delete Data
	// Only Admin can delete the data
	if authClaims.Role != variable.ROLE_ADMIN {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	filter := &model.OrderFilter{
		Id: r.PathValue("id"),
	}

	order, err := h.service.Delete(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, order)
}

func (h *OrderHandler) FindMany(w http.ResponseWriter, r *http.Request) {
	filter := &model.OrderFilter{
		Id:      r.URL.Query().Get("id"),
		Code:    r.URL.Query().Get("code"),
		StoreId: r.URL.Query().Get("store_id"),
	}

	if r.URL.Query().Get("customer.email") != "" {
		filter.Customer = &model.OrderCustomerFilter{
			Email: r.URL.Query().Get("customer.email"),
		}
	}

	// Secure Find Orders
	// Store ID Required
	if filter.StoreId == "" {
		handler.HandleError(w, variable.ErrOrderFilterStoreId)
		return
	}

	authClaims, ok := r.Context().Value(variable.ContextKeyUser).(*model.AuthAccessTokenClaims)
	if !ok {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	storeFilter := &model.StoreFilter{
		Id: filter.StoreId,
	}

	store, err := h.storeService.FindOne(storeFilter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	// Secure Find Orders
	// Only author can view his own store orders
	// Admin can view all store orders
	if !(store.AuthorID == authClaims.UserId || authClaims.Role == variable.ROLE_ADMIN) {
		handler.SendError(w, variable.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	orders, err := h.service.FindMany(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, orders)
}

func (h *OrderHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	filter := &model.OrderFilter{
		Id:      r.URL.Query().Get("id"),
		Code:    r.URL.Query().Get("code"),
		StoreId: r.URL.Query().Get("store_id"),
	}

	if r.URL.Query().Get("customer.email") != "" {
		filter.Customer = &model.OrderCustomerFilter{
			Email: r.URL.Query().Get("customer.email"),
		}
	}

	order, err := h.service.FindOne(filter)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, order)
}
