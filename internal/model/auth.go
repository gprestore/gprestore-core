package model

import "golang.org/x/oauth2"

type Auth struct {
	// Action: LOGIN, REGISTER
	Action string        `json:"action,omitempty"`
	Token  *oauth2.Token `json:"token,omitempty"`
	User   *User         `json:"user,omitempty"`
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
