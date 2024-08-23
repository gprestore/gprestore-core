package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Slug        string             `json:"slug,omitempty" bson:"slug"`
	Name        string             `json:"name,omitempty" bson:"name"`
	Description string             `json:"description,omitempty" bson:"description"`
	AuthorID    string             `json:"author_id,omitempty" bson:"author_id"`
	Logo        string             `json:"logo,omitempty" bson:"logo,omitempty"`
	Banner      string             `json:"banner,omitempty" bson:"banner,omitempty"`
	Badges      []*StoreBadge      `json:"badges" bson:"badges"`
	CreatedAt   *time.Time         `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt   *time.Time         `json:"updated_at,omitempty" bson:"updated_at"`
}

type StoreBadge struct {
	BadgeID   string     `json:"badge_id,omitempty" bson:"badge_id"`
	Badge     *Badge     `json:"badge,omitempty" bson:"-"`
	CreatedAt *time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type StoreCreate struct {
	Slug        string `validate:"required,min=3,lowercase" json:"slug,omitempty" bson:"slug"`
	Name        string `validate:"required,min=3" json:"name,omitempty" bson:"name"`
	Description string `validate:"required,min=8" json:"description,omitempty" bson:"description"`
	AuthorID    string `validate:"required,mongodb" json:"author_id,omitempty" bson:"author_id"`
	Logo        string `validate:"omitempty,url" json:"logo,omitempty" bson:"logo"`
	Banner      string `validate:"omitempty,url" json:"banner,omitempty" bson:"banner"`
}

type StoreUpdate struct {
	Slug        string             `validate:"omitempty,min=3,lowercase" json:"slug,omitempty" bson:"slug"`
	Name        string             `validate:"omitempty,min=3" json:"name,omitempty" bson:"name"`
	Description string             `validate:"omitempty,min=8" json:"description,omitempty" bson:"description"`
	AuthorID    string             `validate:"omitempty,mongodb" json:"author_id,omitempty" bson:"author_id"`
	Logo        string             `validate:"omitempty,url" json:"logo,omitempty" bson:"logo"`
	Banner      string             `validate:"omitempty,url" json:"banner,omitempty" bson:"banner"`
	Badges      []StoreBadgeUpdate `json:"badges,omitempty" bson:"badges"`
}

type StoreBadgeUpdate struct {
	BadgeID string `validate:"omitempty,mongodb" json:"badge_id,omitempty" bson:"badge_id"`
}

type StoreFilter struct {
	Id       string `validate:"omitempty,mongodb" json:"id,omitempty" bson:"_id"`
	Slug     string `validate:"omitempty,min=3,lowercase" json:"slug,omitempty" bson:"slug"`
	AuthorID string `validate:"omitempty,mongodb" json:"author_id,omitempty" bson:"author_id"`
}
