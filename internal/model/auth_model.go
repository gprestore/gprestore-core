package model

type Auth struct {
	// Action: LOGIN, REGISTER
	Action string     `json:"action,omitempty"`
	User   *User      `json:"user,omitempty"`
	Token  *AuthToken `json:"token,omitempty"`
	// Provider: GOOGLE, DISCORD
	Provider string `json:"provider,omitempty"`
}

type AuthToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type AuthAccessTokenClaims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Exp      int64  `json:"exp"`
}

type AuthRefreshTokenClaims struct {
	UserId string `json:"user_id"`
	Exp    int64  `json:"exp"`
}
