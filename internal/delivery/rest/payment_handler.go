package rest

import (
	"net/http"

	"github.com/gprestore/gprestore-core/internal/service"
	"github.com/gprestore/gprestore-core/pkg/handler"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}

func (h *PaymentHandler) FindPaymentChannels(w http.ResponseWriter, r *http.Request) {
	paymentMethods, err := h.service.GetPaymentMethods()
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	handler.SendSuccess(w, paymentMethods)
}
