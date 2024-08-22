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

type UserRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	collection := db.Collection("users")
	_, err := collection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
		{
			Keys:    bson.D{{Key: "phone", Value: 1}},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	return &UserRepository{
		db:         db,
		collection: collection,
	}
}

func (r *UserRepository) Create(input *model.User) (*model.User, error) {
	timeNow := time.Now()
	input.Id = primitive.NewObjectID()
	input.Role = variable.ROLE_USER
	input.CreatedAt = &timeNow
	input.UpdatedAt = &timeNow

	_, err := r.collection.InsertOne(context.TODO(), input)
	return input, err
}

func (r *UserRepository) Update(filter *model.UserFilter, input *model.User) (*model.User, error) {
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

	user, err := r.FindOne(filter)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindMany(filter *model.UserFilter) ([]*model.User, error) {
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

	var users = make([]*model.User, 0)
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindOne(filter *model.UserFilter) (*model.User, error) {
	filterBson, err := converter.InputToBson(filter)
	if err != nil {
		return nil, err
	}

	var user *model.User
	err = r.collection.FindOne(context.TODO(), filterBson).Decode(&user)

	return user, err
}

func (r *UserRepository) Delete(filter *model.UserFilter) (*model.User, error) {
	filterBson, err := converter.InputToBson(filter)
	if err != nil {
		return nil, err
	}

	user, err := r.FindOne(filter)
	if err != nil {
		return nil, err
	}

	_, err = r.collection.DeleteOne(context.TODO(), filterBson)

	return user, err
}
