package database

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQL() *gorm.DB {
	db, err := gorm.Open(postgres.Open(viper.GetString("postgres.url")))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewMongoDB() *mongo.Database {
	ctx := context.TODO()
	mongoDbUrl := viper.GetString("mongodb.url")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbUrl))
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(viper.GetString("mongodb.database"))
}
