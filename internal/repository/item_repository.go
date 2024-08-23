package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gosimple/slug"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/pkg/converter"
	"github.com/gprestore/gprestore-core/pkg/random"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ItemRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewItemRepository(db *mongo.Database) *ItemRepository {
	collection := db.Collection("items")
	_, err := collection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "slug", Value: 1}},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
		{
			Keys:    bson.D{{Key: "store_id", Value: 1}},
			Options: options.Index(),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	return &ItemRepository{
		db:         db,
		collection: collection,
	}
}

func (r *ItemRepository) Create(input *model.Item) (*model.Item, error) {
	timeNow := time.Now()
	input.Id = primitive.NewObjectID()
	input.Slug = slug.Make(input.Name + " " + random.Number(5))
	input.Categories = make([]model.ItemCategory, 0)
	input.CreatedAt = &timeNow
	input.UpdatedAt = &timeNow

	_, err := r.collection.InsertOne(context.TODO(), input)
	return input, err
}

func (r *ItemRepository) Update(filter *model.ItemFilter, input *model.Item) (*model.Item, error) {
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

	item, err := r.FindOne(filter)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (r *ItemRepository) FindMany(filter *model.ItemFilter) ([]*model.Item, error) {
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

	var stores = make([]*model.Item, 0)
	err = cursor.All(context.TODO(), &stores)
	if err != nil {
		return nil, err
	}

	return stores, nil
}

func (r *ItemRepository) FindOne(filter *model.ItemFilter) (*model.Item, error) {
	filterBson, err := converter.InputToBson(filter)
	if err != nil {
		return nil, err
	}

	var item *model.Item
	err = r.collection.FindOne(context.TODO(), filterBson).Decode(&item)

	return item, err
}

func (r *ItemRepository) Delete(filter *model.ItemFilter) (*model.Item, error) {
	filterBson, err := converter.InputToBson(filter)
	if err != nil {
		return nil, err
	}

	item, err := r.FindOne(filter)
	if err != nil {
		return nil, err
	}

	_, err = r.collection.DeleteOne(context.TODO(), filterBson)

	return item, err
}
