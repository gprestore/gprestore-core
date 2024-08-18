package user

import (
	"context"
	"log"
	"time"

	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/pkg/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
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

	return &Repository{
		db:         db,
		collection: collection,
	}
}

func (r *Repository) Create(input *model.User) (*model.User, error) {
	timeNow := time.Now()
	input.Id = primitive.NewObjectID()
	input.CreatedAt = &timeNow
	input.UpdatedAt = &timeNow

	result, err := r.collection.InsertOne(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	input.Id = result.InsertedID.(primitive.ObjectID)

	return input, nil
}

func (r *Repository) Update(id string, input *model.User) (*model.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()
	input.UpdatedAt = &timeNow

	inputBson, err := converter.InputToBson(input)
	if err != nil {
		log.Println("Errrrr")
		return nil, err
	}

	result, err := r.collection.UpdateByID(context.TODO(), objectId, bson.D{
		{
			Key:   "$set",
			Value: inputBson,
		},
	})
	if err != nil {
		return nil, err
	}
	log.Println(result)
	return input, nil
}

func (r *Repository) FindMany() ([]*model.User, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var users []*model.User
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) FindOne(id string) (*model.User, error) {
	return nil, nil
}

func (r *Repository) Delete(id string) (*model.User, error) {
	return nil, nil
}
