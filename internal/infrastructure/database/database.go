package database

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB() *mongo.Database {
	ctx := context.TODO()
	mongoDbUrl := viper.GetString("mongodb.url")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbUrl))
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(viper.GetString("mongodb.database"))
}
