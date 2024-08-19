package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Username string             `json:"username,omitempty" bson:"username"`
	FullName string             `json:"fullName,omitempty" bson:"full_name"`
	Email    string             `json:"email,omitempty" bson:"email"`
	Phone    string             `json:"phone,omitempty" bson:"phone"`
	// Role: USER, SELLER, ADMIN
	Role         string           `json:"role,omitempty" bson:"role"`
	VerifyStatus UserVerifyStatus `json:"verify_status,omitempty" bson:"verify_status"`
	CreatedAt    *time.Time       `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt    *time.Time       `json:"updatedAt,omitempty" bson:"updated_at"`
}

type UserVerifyStatus struct {
	Email bool `json:"email,omitempty" bson:"email"`
	Phone bool `json:"phone,omitempty" bson:"phone"`
}

type UserCreate struct {
	Username string `validate:"required,min=3" bson:"username" json:"username,omitempty"`
	FullName string `validate:"required,min=3" bson:"full_name" json:"full_name,omitempty"`
	Email    string `validate:"required,email" bson:"email" json:"email,omitempty"`
	Phone    string `validate:"required,e164" bson:"phone" json:"phone,omitempty"`
}

type UserUpdate struct {
	Username     string           `validate:"omitempty,min=3" bson:"username" json:"username,omitempty"`
	FullName     string           `validate:"omitempty,min=3" bson:"full_name" json:"full_name,omitempty"`
	Email        string           `validate:"omitempty,email" bson:"email" json:"email,omitempty"`
	Phone        string           `validate:"omitempty,e164" bson:"phone" json:"phone,omitempty"`
	VerifyStatus UserVerifyStatus `bson:"verify_status" json:"verify_status,omitempty"`
}

type UserFilter struct {
	Id       string `validate:"omitempty,uuid" bson:"_id" json:"id,omitempty"`
	Username string `validate:"omitempty,min=3" bson:"username" json:"username,omitempty"`
	Email    string `validate:"omitempty,email" bson:"email" json:"email,omitempty"`
	Phone    string `validate:"omitempty,e164" bson:"phone" json:"phone,omitempty"`
}
