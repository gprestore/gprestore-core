package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username"`
	FullName string             `json:"full_name,omitempty" bson:"full_name"`
	Email    string             `json:"email,omitempty" bson:"email"`
	Phone    *string            `json:"phone,omitempty" bson:"phone"`
	// Role: USER, SELLER, ADMIN
	Role         string           `json:"role,omitempty" bson:"role"`
	VerifyStatus UserVerifyStatus `json:"verify_status,omitempty" bson:"verify_status"`
	Image        *string          `json:"image,omitempty"`
	CreatedAt    *time.Time       `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    *time.Time       `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UserVerifyStatus struct {
	Email bool `json:"email" bson:"email"`
	Phone bool `json:"phone" bson:"phone"`
}

type UserCreate struct {
	Username     string           `validate:"required,min=3" bson:"username,omitempty" json:"username,omitempty"`
	FullName     string           `validate:"required,min=3" bson:"full_name,omitempty" json:"full_name,omitempty"`
	Email        string           `validate:"required,email" bson:"email,omitempty" json:"email,omitempty"`
	Phone        *string          `validate:"omitempty,e164" bson:"phone,omitempty" json:"phone,omitempty"`
	Image        string           `validate:"omitempty,url" bson:"image,omitempty" json:"image,omitempty"`
	VerifyStatus UserVerifyStatus `bson:"verify_status,omitempty" json:"verify_status,omitempty"`
}

type UserUpdate struct {
	Username     string           `validate:"omitempty,min=3" bson:"username,omitempty" json:"username,omitempty"`
	FullName     string           `validate:"omitempty,min=3" bson:"full_name,omitempty" json:"full_name,omitempty"`
	Email        string           `validate:"omitempty,email" bson:"email,omitempty" json:"email,omitempty"`
	Phone        string           `validate:"omitempty,e164" bson:"phone,omitempty" json:"phone,omitempty"`
	Image        *string          `validate:"omitempty,url" bson:"image,omitempty" json:"image,omitempty"`
	VerifyStatus UserVerifyStatus `bson:"verify_status,omitempty" json:"verify_status,omitempty"`
}

type UserFilter struct {
	Id       string `validate:"omitempty,mongodb" bson:"_id,omitempty" json:"id,omitempty"`
	Username string `validate:"omitempty,min=3" bson:"username,omitempty" json:"username,omitempty"`
	Email    string `validate:"omitempty,email" bson:"email,omitempty" json:"email,omitempty"`
	Phone    string `validate:"omitempty,e164" bson:"phone,omitempty" json:"phone,omitempty"`
}
