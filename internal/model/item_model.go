package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	StoreId     string             `json:"store_id,omitempty" bson:"store_id,omitempty"`
	Slug        string             `json:"slug,omitempty" bson:"slug,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       int                `json:"price,omitempty" bson:"price,omitempty"`
	Categories  []ItemCategory     `json:"categories" bson:"categories"`
	StockCount  *int               `json:"stock_count" bson:"stock_count,omitempty"`
	CreatedAt   *time.Time         `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   *time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type ItemCategory struct {
	CategoryID string    `json:"category_id,omitempty" bson:"category_id,omitempty"`
	Category   *Category `json:"category,omitempty" bson:"-"`
}

type ItemCreate struct {
	StoreId     string         `validate:"required,mongodb" json:"store_id,omitempty" bson:"store_id,omitempty"`
	Name        string         `validate:"required,min=3" json:"name,omitempty" bson:"name,omitempty"`
	Description string         `validate:"required,min=8" json:"description,omitempty" bson:"description,omitempty"`
	Categories  []ItemCategory `json:"categories,omitempty" bson:"categories,omitempty"`
	Price       int            `validate:"required,min=1000" json:"price,omitempty" bson:"price,omitempty"`
}

type ItemUpdate struct {
	StoreId     string         `validate:"omitempty,mongodb" json:"store_id,omitempty" bson:"store_id,omitempty"`
	Slug        string         `validate:"omitempty,min=3,lowercase" json:"slug,omitempty" bson:"slug,omitempty"`
	Name        string         `validate:"omitempty,min=3" json:"name,omitempty" bson:"name,omitempty"`
	Description string         `validate:"omitempty,min=8" json:"description,omitempty" bson:"description,omitempty"`
	Categories  []ItemCategory `json:"categories,omitempty" bson:"categories,omitempty"`
	Price       int            `validate:"omitempty,min=1000" json:"price,omitempty" bson:"price,omitempty"`
}

type ItemFilter struct {
	Id      string `validate:"omitempty,mongodb" json:"id,omitempty" bson:"_id,omitempty"`
	StoreId string `validate:"omitempty,mongodb" json:"store_id,omitempty" bson:"store_id,omitempty"`
	Slug    string `validate:"omitempty,min=3,lowercase" json:"slug,omitempty" bson:"slug,omitempty"`
}
