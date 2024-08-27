package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AWAITING_PAYMENT, PAYMENT_SUCCESS, COMPLETED, EXPIRED
type OrderStatus string

// flip.FlipCode...
type PaymentBankCode string

// flip.FlipAccountType...
type PaymentBankType string

type Order struct {
	Id              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Code            string             `json:"code,omitempty" bson:"code,omitempty"`
	StoreId         string             `json:"store_id,omitempty" bson:"store_id,omitempty"`
	Items           []OrderItem        `json:"items,omitempty" bson:"items,omitempty"`
	PaymentBankType PaymentBankType    `json:"payment_bank_type,omitempty" bson:"payment_bank_type,omitempty"`
	PaymentBankCode PaymentBankCode    `json:"payment_bank_code,omitempty" bson:"payment_bank_code,omitempty"`
	Payment         *Payment           `json:"payment,omitempty" bson:"payment,omitempty"`
	Fees            []OrderFee         `json:"fees,omitempty" bson:"fees,omitempty"`
	Customer        OrderCustomer      `json:"customer,omitempty" bson:"customer,omitempty"`
	Subtotal        int                `json:"subtotal,omitempty" bson:"subtotal,omitempty"`
	Status          OrderStatus        `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt       *time.Time         `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt       *time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type OrderItem struct {
	ItemId   string `validate:"required,mongodb" json:"item_id,omitempty" bson:"item_id,omitempty"`
	Name     string `validate:"omitempty,min=3" json:"name,omitempty" bson:"name,omitempty"`
	Price    int    `validate:"omitempty" json:"price,omitempty" bson:"price,omitempty"`
	Quantity int    `validate:"required,min=1" json:"quantity,omitempty" bson:"quantity,omitempty"`
}

type OrderCustomer struct {
	Name  string `validate:"required,min=3" json:"name,omitempty" bson:"name,omitempty"`
	Email string `validate:"required,email" json:"email,omitempty" bson:"email,omitempty"`
}

type OrderPayment struct {
	Id            string
	Reference     string
	Channel       string
	AccountNumber string
	QrisString    string
	Status        string
}

type OrderFee struct {
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	Amount int    `json:"amount,omitempty" bson:"amount,omitempty"`
}

type OrderCreate struct {
	StoreId         string          `validate:"required,mongodb" json:"store_id,omitempty" bson:"store_id,omitempty"`
	Items           []OrderItem     `validate:"required,min=1" json:"items,omitempty" bson:"items,omitempty"`
	Customer        OrderCustomer   `validate:"required" json:"customer,omitempty" bson:"customer,omitempty"`
	PaymentBankType PaymentBankType `validate:"required" json:"payment_bank_type,omitempty" bson:"payment_bank_type,omitempty"`
	PaymentBankCode PaymentBankCode `validate:"required" json:"payment_bank_code,omitempty" bson:"payment_bank_code,omitempty"`
}

type OrderUpdate struct {
	Status OrderStatus `json:"status,omitempty" bson:"status,omitempty"`
}

type OrderFilter struct {
	Id       string         `validate:"omitempty,mongodb" json:"id,omitempty" bson:"_id,omitempty"`
	Code     string         `validate:"omitempty" json:"code,omitempty" bson:"code,omitempty"`
	StoreId  string         `validate:"omitempty,mongodb" json:"store_id,omitempty" bson:"store_id,omitempty"`
	Customer *OrderCustomer `validate:"omitempty" json:"customer,omitempty" bson:"customer,omitempty"`
}
