package service

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/pkg/flip"
)

type PaymentService struct {
	flipClient *flip.FlipClient
	validate   *validator.Validate
}

func NewPaymentService(flipClient *flip.FlipClient, validate *validator.Validate) *PaymentService {
	return &PaymentService{
		flipClient: flipClient,
		validate:   validate,
	}
}

func (s *PaymentService) Create(order *model.Order) (*model.Payment, error) {
	request := &flip.FlipBillRequest{
		Title:          order.Code,
		Type:           flip.FlipBillTypeSingle,
		Amount:         order.Subtotal,
		Step:           "3",
		SenderName:     order.Customer.Name,
		SenderEmail:    order.Customer.Email,
		SenderBankType: string(order.PaymentBankType),
		SenderBank:     string(order.PaymentBankCode),
	}

	bill, err := s.flipClient.CreatePayment(request)
	if err != nil {
		return nil, err
	}

	payment := &model.Payment{
		Id:     bill.BillPayment.Id,
		LinkId: fmt.Sprintf("%d", bill.LinkId),
		Status: model.PaymentStatus(bill.Status),
		PaymentChannel: &model.PaymentChannel{
			AccountNumber: bill.BillPayment.ReceiverBankAccount.AccountNumber,
			AccountType:   bill.BillPayment.ReceiverBankAccount.AccountType,
			BankCode:      bill.BillPayment.ReceiverBankAccount.BankCode,
			AccountHolder: bill.BillPayment.ReceiverBankAccount.AccountHolder,
			QrCodeData:    bill.BillPayment.ReceiverBankAccount.QrCodeData,
		},
	}

	return payment, nil
}
