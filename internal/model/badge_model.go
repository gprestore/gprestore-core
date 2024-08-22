package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Badge struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Slug        string             `json:"slug,omitempty" bson:"slug,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Icon        *string            `json:"icon,omitempty" bson:"icon,omitempty"`
	CreatedAt   *time.Time         `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   *time.Time         `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
