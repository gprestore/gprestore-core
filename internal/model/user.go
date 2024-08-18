package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id,omitempty"`
	Username  string    `gorm:"unique" json:"username,omitempty"`
	FullName  string    `json:"full_name,omitempty"`
	Email     string    `gorm:"unique" json:"email,omitempty"`
	Phone     string    `gorm:"unique" json:"phone,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}

type UserCreateValidation struct {
	Username string `validate:"required,min=3"`
	FullName string `validate:"required,min=3"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,e164"`
}

type UserUpdateValidation struct {
	Username string `validate:"omitempty,min=3"`
	FullName string `validate:"omitempty,min=3"`
	Email    string `validate:"omitempty,email"`
	Phone    string `validate:"omitempty,e164"`
}

type UserFindValidation struct {
	Id       string `validate:"omitempty,uuid"`
	Username string `validate:"omitempty,min=3"`
	FullName string `validate:"omitempty,min=3"`
	Email    string `validate:"omitempty,email"`
	Phone    string `validate:"omitempty,e164"`
}
