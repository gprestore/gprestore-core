package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/gprestore/gprestore-core/internal/config"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var paymentService *service.PaymentService

func init() {
	config.Load()
	paymentService = service.NewPaymentService()
}

func TestPayment(t *testing.T) {
	order := &model.Order{
		Code:     "GPR-0123",
		Subtotal: 10000,
	}
	inv, err := paymentService.CreateInvoice(order)
	log.Println("Error:")
	log.Println(err)
	log.Println(inv.InvoiceUrl)
}

func TestGetPaymentMethods(t *testing.T) {
	methods, err := paymentService.GetPaymentMethods()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(methods)
}

func TestCreatePayment(t *testing.T) {
	order := &model.Order{
		Id:      primitive.NewObjectID(),
		StoreId: primitive.NewObjectID().Hex(),
		Items: []model.OrderItem{
			{
				ItemId:   primitive.NilObjectID.Hex(),
				Name:     "Minecraft Alts",
				Price:    10000,
				Quantity: 1,
			},
		},
		Code:     "GPR-0123",
		Subtotal: 10000,
		Customer: model.OrderCustomer{
			Name:  "Agil Ghani Istikmal",
			Email: "agil_g@safatanc.com",
		},
	}
	resp, err := paymentService.CreatePayment(order)
	if err != nil {
		log.Fatal(err)
	}
	order.PaymentChannel = resp

	fmt.Println(*order.PaymentChannel.QrCode.Get().ChannelProperties.QrString)
}
