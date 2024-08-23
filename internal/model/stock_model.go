package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Stock struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	StoreId   string             `json:"store_id,omitempty" bson:"store_id,omitempty"`
	ItemId    string             `json:"item_id,omitempty" bson:"item_id,omitempty"`
	Separator string             `json:"separator,omitempty" bson:"separator,omitempty"`
	Contents  []string           `json:"contents,omitempty" bson:"contents,omitempty"`
	Count     int                `json:"count,omitempty" bson:"count,omitempty"`
	CreatedAt *time.Time         `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt *time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type StockCreate struct {
	StoreId   string   `validate:"required,mongodb" json:"store_id,omitempty" bson:"store_id,omitempty"`
	ItemId    string   `validate:"required,mongodb" json:"item_id,omitempty" bson:"item_id,omitempty"`
	Separator string   `validate:"required" json:"separator,omitempty" bson:"separator,omitempty"`
	Contents  []string `json:"contents,omitempty" bson:"contents,omitempty"`
}

type StockUpdate struct {
	StoreId   string   `validate:"omitempty,mongodb" json:"store_id,omitempty" bson:"store_id,omitempty"`
	ItemId    string   `validate:"omitempty,mongodb" json:"item_id,omitempty" bson:"item_id,omitempty"`
	Separator string   `validate:"omitempty" json:"separator,omitempty" bson:"separator,omitempty"`
	Contents  []string `json:"contents,omitempty" bson:"contents,omitempty"`
}

type StockFilter struct {
	Id        string `json:"id,omitempty" bson:"_id,omitempty"`
	StoreId   string `validate:"omitempty,mongodb" json:"store_id,omitempty" bson:"store_id,omitempty"`
	ItemId    string `validate:"omitempty,mongodb" json:"item_id,omitempty" bson:"item_id,omitempty"`
	Separator string `validate:"omitempty" json:"separator,omitempty" bson:"separator,omitempty"`
}
