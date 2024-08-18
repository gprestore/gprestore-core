package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MongoDBUrl string
}

var AppConfig *Config

func Load() {
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
