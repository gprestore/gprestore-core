package config

import (
	"log"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	OAuthGoogle oauth2.Config
}

var AppConfig Config

func Load() {
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	googleConfig()
}

func googleConfig() {
	AppConfig.OAuthGoogle = oauth2.Config{
		RedirectURL:  viper.GetString("oauth.callback_url"),
		ClientID:     viper.GetString("oauth.google.client_id"),
		ClientSecret: viper.GetString("oauth.google.client_secret"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
