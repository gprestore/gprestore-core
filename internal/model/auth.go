package model

import (
	"time"
)

type Auth struct {
	// Action: LOGIN, REGISTER
	Action string     `json:"action,omitempty"`
	User   *User      `json:"user,omitempty"`
	Token  *AuthToken `json:"token,omitempty"`
	// Provider: GOOGLE, DISCORD
	Provider string `json:"provider,omitempty"`
}

type AuthToken struct {
	AccessToken  string     `json:"access_token,omitempty"`
	RefreshToken string     `json:"refresh_token,omitempty"`
	ExpiryAt     *time.Time `json:"expiry_at,omitempty"`
}

type AuthGoogle struct {
	Id            string `json:"id,omitempty"`
	Email         string `json:"email,omitempty"`
	VerifiedEmail bool   `json:"verified_email,omitempty"`
	Name          string `json:"name,omitempty"`
	GivenName     string `json:"given_name,omitempty"`
	FamilyName    string `json:"family_name,omitempty"`
	Picture       string `json:"picture,omitempty"`
	Locale        string `json:"locale,omitempty"`
}
