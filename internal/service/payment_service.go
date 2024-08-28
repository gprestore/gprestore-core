package service

import (
	"fmt"
	"strconv"

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

func (s *PaymentService) GetBill(linkId string) (*flip.FlipBill, error) {
	id, _ := strconv.Atoi(linkId)
	bill, err := s.flipClient.GetBill(id)
	if err != nil {
		return nil, err
	}

	return bill, err
}

func (s *PaymentService) GetPayment(linkId string) (*flip.FlipPayment, error) {
	id, _ := strconv.Atoi(linkId)
	payment, err := s.flipClient.GetPayment(id)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *PaymentService) GetPaymentData(linkId string) (*flip.FlipPaymentData, error) {
	payment, err := s.GetPayment(linkId)
	if err != nil {
		return nil, err
	}

	if len(payment.Data) == 0 {
		return nil, fmt.Errorf("no document in result")
	}

	return &payment.Data[0], nil
}
