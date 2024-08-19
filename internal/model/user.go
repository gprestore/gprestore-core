package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Username  string             `json:"username,omitempty" bson:"username"`
	FullName  string             `json:"fullName,omitempty" bson:"full_name"`
	Email     string             `json:"email,omitempty" bson:"email"`
	Phone     string             `json:"phone,omitempty" bson:"phone"`
	CreatedAt *time.Time         `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt *time.Time         `json:"updatedAt,omitempty" bson:"updated_at"`
}

type UserCreate struct {
	Username string `validate:"required,min=3"`
	FullName string `validate:"required,min=3"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,e164"`
}

type UserUpdate struct {
	Username string `validate:"omitempty,min=3"`
	FullName string `validate:"omitempty,min=3"`
	Email    string `validate:"omitempty,email"`
	Phone    string `validate:"omitempty,e164"`
}

type UserFilter struct {
	Id       string `validate:"omitempty,uuid"`
	Username string `validate:"omitempty,min=3"`
	Email    string `validate:"omitempty,email"`
	Phone    string `validate:"omitempty,e164"`
}
