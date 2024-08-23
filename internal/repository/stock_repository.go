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

type StockRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewStockRepository(db *mongo.Database) *StockRepository {
	collection := db.Collection("stocks")
	_, err := collection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "item_id", Value: 1}},
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

	return &StockRepository{
		db:         db,
		collection: collection,
	}
}

func (r *StockRepository) Create(input *model.Stock) (*model.Stock, error) {
	timeNow := time.Now()
	input.Id = primitive.NewObjectID()
	input.Contents = make([]string, 0)
	input.CreatedAt = &timeNow
	input.UpdatedAt = &timeNow

	_, err := r.collection.InsertOne(context.TODO(), input)
	return input, err
}

func (r *StockRepository) Update(filter *model.StockFilter, input *model.Stock) (*model.Stock, error) {
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

	stock, err := r.FindOne(filter)
	if err != nil {
		return nil, err
	}

	return stock, nil
}

func (r *StockRepository) FindMany(filter *model.StockFilter) ([]*model.Stock, error) {
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

	var stores = make([]*model.Stock, 0)
	err = cursor.All(context.TODO(), &stores)
	if err != nil {
		return nil, err
	}

	return stores, nil
}

func (r *StockRepository) FindOne(filter *model.StockFilter) (*model.Stock, error) {
	filterBson, err := converter.InputToBson(filter)
	if err != nil {
		return nil, err
	}

	var stock *model.Stock
	err = r.collection.FindOne(context.TODO(), filterBson).Decode(&stock)

	return stock, err
}

func (r *StockRepository) Delete(filter *model.StockFilter) (*model.Stock, error) {
	filterBson, err := converter.InputToBson(filter)
	if err != nil {
		return nil, err
	}

	stock, err := r.FindOne(filter)
	if err != nil {
		return nil, err
	}

	_, err = r.collection.DeleteOne(context.TODO(), filterBson)

	return stock, err
}
