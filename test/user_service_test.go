package test

import (
	"log"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gprestore/gprestore-core/internal/config"
	"github.com/gprestore/gprestore-core/internal/database"
	"github.com/gprestore/gprestore-core/internal/domain/user"
	"github.com/gprestore/gprestore-core/internal/model"
)

var s *user.Service

func init() {
	config.Load()

	db := database.NewMongoDB()
	v := validator.New()
	r := user.NewRepository(db)
	s = user.NewService(r, v)
}

func TestCreateUserService(t *testing.T) {
	input := &model.UserCreate{
		Username: "nabilmz",
		FullName: "Nabil Meizahir",
		Email:    "nabil_mz@safatanc.com",
		Phone:    "+6281234567890",
	}
	user, err := s.Create(input)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
}

func TestUpdateUserService(t *testing.T) {
	filter := &model.UserFilter{
		Username: "agilistikmal",
	}
	input := &model.UserUpdate{
		VerifyStatus: model.UserVerifyStatus{
			Email: true,
		},
		Phone: "+62812345678",
	}
	user, err := s.Update(filter, input)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)
}

func TestFindUsersService(t *testing.T) {
	filter := &model.UserFilter{
		Username: "agilistikmal",
	}

	user, err := s.FindOne(filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
}
