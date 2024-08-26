package service

import (
	"context"

	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/spf13/viper"
	"github.com/xendit/xendit-go/v6"
	"github.com/xendit/xendit-go/v6/invoice"
	"github.com/xendit/xendit-go/v6/payment_method"
	"github.com/xendit/xendit-go/v6/payment_request"
)

type PaymentService struct {
	xenditClient *xendit.APIClient
}

func NewPaymentService() *PaymentService {
	client := xendit.NewClient(viper.GetString("xendit.secret_key"))

	return &PaymentService{
		xenditClient: client,
	}
}

func (s *PaymentService) CreateInvoice(order *model.Order) (*invoice.Invoice, error) {
	request := *invoice.NewCreateInvoiceRequest(order.Code, float64(order.Subtotal))
	resp, _, err := s.xenditClient.InvoiceApi.
		CreateInvoice(context.Background()).
		CreateInvoiceRequest(request).
		Execute()

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PaymentService) CreatePayment(order *model.Order) (*payment_request.PaymentMethod, error) {
	amount := float64(order.Subtotal)
	var xenditItems []payment_request.PaymentRequestBasketItem
	for _, item := range order.Items {
		xenditItems = append(xenditItems, payment_request.PaymentRequestBasketItem{
			ReferenceId: &item.ItemId,
			Name:        item.Name,
			Type:        &order.StoreId,
			Category:    order.StoreId,
			Currency:    "IDR",
			Price:       float64(item.Price),
			Quantity:    float64(item.Quantity),
		})
	}

	request := payment_request.PaymentRequestParameters{
		ReferenceId: &order.Code,
		Amount:      &amount,
		Currency:    payment_request.PAYMENTREQUESTCURRENCY_IDR,
		Items:       xenditItems,
		Metadata: map[string]interface{}{
			"name":  order.Customer.Name,
			"email": order.Customer.Email,
		},
		PaymentMethod: payment_request.NewPaymentMethodParameters(
			payment_request.PAYMENTMETHODTYPE_QR_CODE,
			payment_request.PAYMENTMETHODREUSABILITY_ONE_TIME_USE,
		),
	}

	resp, _, err := s.xenditClient.PaymentRequestApi.CreatePaymentRequest(context.Background()).PaymentRequestParameters(request).Execute()
	if err != nil {
		return nil, err
	}

	return &resp.PaymentMethod, nil
}

func (s *PaymentService) GetPaymentMethods() ([]payment_method.PaymentMethod, error) {
	resp, _, err := s.xenditClient.PaymentMethodApi.GetAllPaymentMethods(context.Background()).Execute()
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}
