package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Slug        string             `json:"slug,omitempty" bson:"slug,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
}
