package test

import (
	"log"
	"testing"

	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/pkg/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInputToBson(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66c1ebd717efaae92bb407f0")
	if err != nil {
		log.Fatal(err)
	}
	input := &model.User{
		Id:    id,
		Email: "asdas@gmail.com",
	}
	result, err := converter.InputToBson(input)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}
