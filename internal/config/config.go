package config

import (
	"log"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/google"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
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

	oauthConfig()
}

func oauthConfig() {
	goth.UseProviders(
		discord.New(
			viper.GetString("oauth.discord.client_id"),
			viper.GetString("oauth.discord.client_secret"),
			viper.GetString("oauth.discord.callback_url"),
			discord.ScopeEmail,
			discord.ScopeIdentify,
			discord.ScopeJoinGuild,
		),
		google.New(
			viper.GetString("oauth.google.client_id"),
			viper.GetString("oauth.google.client_secret"),
			viper.GetString("oauth.google.callback_url"),
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		),
	)
}
