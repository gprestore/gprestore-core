package service

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/repository"
	"github.com/gprestore/gprestore-core/pkg/converter"
	"github.com/gprestore/gprestore-core/pkg/variable"
	"github.com/markbates/goth"
	"github.com/spf13/viper"
)

type AuthService struct {
	userRepository *repository.UserRepository
	validate       *validator.Validate
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) NewAccessWithRefreshToken(user *model.User) (*model.AuthToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.Id.Hex(),
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(1 * time.Hour).Unix(),
	})

	accessToken, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return nil, err
	}

	return s.NewRefreshToken(accessToken)
}

func (s *AuthService) NewRefreshToken(accessToken string) (*model.AuthToken, error) {
	expiryAt := time.Now().Add(24 * 14 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": accessToken,
		"exp":   expiryAt,
	})

	refreshToken, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return nil, err
	}

	return &model.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiryAt:     &expiryAt,
	}, nil
}

func (s *AuthService) FindUser(filter *model.UserFilter) (*model.User, error) {
	user, err := s.userRepository.FindOne(filter)
	return user, err
}

func (s *AuthService) CreateUser(input *model.UserCreate) (*model.User, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	inputUser, err := converter.StructConverter[model.User](input)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.Create(inputUser)
	return user, err
}

func (s *AuthService) LoginOrRegister(gothUser *goth.User) (*model.Auth, error) {
	auth := &model.Auth{
		Action:   variable.AUTH_ACTION_LOGIN,
		Provider: gothUser.Provider,
	}

	filter := &model.UserFilter{
		Email: gothUser.Email,
	}

	user, err := s.FindUser(filter)
	if err != nil {
		auth.Action = variable.AUTH_ACTION_REGISTER

		input := &model.UserCreate{
			Username: "user" + gothUser.UserID,
			FullName: gothUser.Name,
			Email:    gothUser.Email,
			VerifyStatus: model.UserVerifyStatus{
				Email: true,
			},
			Image: gothUser.AvatarURL,
		}

		user, err = s.CreateUser(input)
		if err != nil {
			return nil, err
		}
	}

	authToken, err := s.NewAccessWithRefreshToken(user)
	if err != nil {
		return nil, err
	}

	auth.User = user
	auth.Token = authToken

	return auth, nil
}

func (s *AuthService) ValidateAccessToken(accessToken string) (*model.User, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, fmt.Errorf("invalid token")
	}

	filter := &model.UserFilter{
		Id: claims["user_id"].(string),
	}

	user, err := s.userRepository.FindOne(filter)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) ValidateRefreshToken(refreshtoken string) (*model.User, error) {
	token, err := jwt.Parse(refreshtoken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, fmt.Errorf("invalid token")
	}

	parser := jwt.Parser{}
	token, _, err = parser.ParseUnverified(claims["token"].(string), jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	claims, ok = token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, fmt.Errorf("invalid token")
	}

	filter := &model.UserFilter{
		Id: claims["user_id"].(string),
	}
	user, err := s.userRepository.FindOne(filter)
	if err != nil {
		return nil, err
	}
	return user, nil
}
