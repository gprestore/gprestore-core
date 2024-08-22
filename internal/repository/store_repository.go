package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/pkg/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *StoreRepository) Create(input *model.Store) (*model.Store, error) {
	timeNow := time.Now()
	input.Id = primitive.NewObjectID()
	input.Badges = make([]model.StoreBadge, 0)
	input.CreatedAt = &timeNow
	input.UpdatedAt = &timeNow

	_, err := r.collection.InsertOne(context.TODO(), input)
	return input, err
}

func (r *StoreRepository) Update(filter *model.StoreFilter, input *model.Store) (*model.Store, error) {
	filterBson, err := converter.InputToBson(filter)
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()
	input.UpdatedAt = &timeNow

	inputBson, err := converter.InputToBson(input)
	if err != nil {
		return nil, err
	}

	result, err := r.collection.UpdateOne(context.TODO(), filterBson, bson.D{
		{
			Key:   "$set",
			Value: inputBson,
		},
	})
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("mongo: no documents in result")
	}

	store, err := r.FindOne(filter)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (r *StoreRepository) FindMany(filter *model.StoreFilter) ([]*model.Store, error) {
	filterBson := bson.D{}
	if filter != nil {
		fb, err := converter.InputToBson(filter)
		if err != nil {
			return nil, err
		}
		filterBson = fb
	}

	cursor, err := r.collection.Find(context.TODO(), filterBson)
	if err != nil {
		return nil, err
	}

	var stores = make([]*model.Store, 0)
	err = cursor.All(context.TODO(), &stores)
	if err != nil {
		return nil, err
	}

	return stores, nil
}

func (r *StoreRepository) FindOne(filter *model.StoreFilter) (*model.Store, error) {
	filterBson, err := converter.InputToBson(filter)
	if err != nil {
		return nil, err
	}

	var store *model.Store
	err = r.collection.FindOne(context.TODO(), filterBson).Decode(&store)

	return store, err
}

func (r *StoreRepository) Delete(filter *model.StoreFilter) (*model.Store, error) {
	filterBson, err := converter.InputToBson(filter)
	if err != nil {
		return nil, err
	}

	store, err := r.FindOne(filter)
	if err != nil {
		return nil, err
	}

	_, err = r.collection.DeleteOne(context.TODO(), filterBson)

	return store, err
}
