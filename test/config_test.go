package test

import (
	"log"
	"testing"

	"github.com/gprestore/gprestore-core/internal/config"
	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	config.Load()
	mongoDbUrl := viper.GetString("MONGODB_URL")
	log.Println(mongoDbUrl)
}

func TestConfigMap(t *testing.T) {
	config.Load()
	accounts := viper.Get("mail.accounts")

	for _, account := range accounts.([]any) {
		username := account.(map[string]any)["username"].(string)
		password := account.(map[string]any)["password"].(string)

		log.Println(username, password)
	}
}
