package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/pkg/converter"
	"github.com/gprestore/gprestore-core/pkg/variable"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	collection := db.Collection("orders")
	_, err := collection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "code", Value: 1}},
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

	return &OrderRepository{
		db:         db,
		collection: collection,
	}
}

func (r *OrderRepository) Create(input *model.Order) (*model.Order, error) {
	timeNow := time.Now()
	input.Id = primitive.NewObjectID()
	input.Status = variable.ORDER_AWAITING_PAYMENT
	input.CreatedAt = &timeNow
	input.UpdatedAt = &timeNow

	_, err := r.collection.InsertOne(context.TODO(), input)
	return input, err
}

func (r *OrderRepository) Update(filter *model.OrderFilter, input *model.Order) (*model.Order, error) {
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

	order, err := r.FindOne(filter)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) FindMany(filter *model.OrderFilter) ([]*model.Order, error) {
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

	var stores = make([]*model.Order, 0)
	err = cursor.All(context.TODO(), &stores)
	if err != nil {
		return nil, err
	}

	return stores, nil
}

func (r *OrderRepository) FindOne(filter *model.OrderFilter) (*model.Order, error) {
	filterBson, err := converter.InputToBson(filter)
	if err != nil {
		return nil, err
	}

	var order *model.Order
	err = r.collection.FindOne(context.TODO(), filterBson).Decode(&order)

	return order, err
}

func (r *OrderRepository) Delete(filter *model.OrderFilter) (*model.Order, error) {
	filterBson, err := converter.InputToBson(filter)
	if err != nil {
		return nil, err
	}

	order, err := r.FindOne(filter)
	if err != nil {
		return nil, err
	}

	_, err = r.collection.DeleteOne(context.TODO(), filterBson)

	return order, err
}
