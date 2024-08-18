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
