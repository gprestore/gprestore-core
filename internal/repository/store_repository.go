package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StoreRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewStoreRepository(db *mongo.Database) *StoreRepository {
	collection := db.Collection("stores")
	_, err := collection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "slug", Value: 1}},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
		{
			Keys:    bson.D{{Key: "author_id", Value: 1}},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	return &StoreRepository{
		db:         db,
		collection: collection,
	}
}

func (r *StoreRepository) Create()
