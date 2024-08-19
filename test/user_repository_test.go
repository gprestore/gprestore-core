package test

import (
	"log"
	"testing"

	"github.com/gprestore/gprestore-core/internal/config"
	"github.com/gprestore/gprestore-core/internal/database"
	"github.com/gprestore/gprestore-core/internal/domain/user"
	"github.com/gprestore/gprestore-core/internal/model"
)

var r *user.Repository

func init() {
	config.Load()

	db := database.NewMongoDB()
	r = user.NewRepository(db)
}

func TestCreateUserRepository(t *testing.T) {
	user, err := r.Create(&model.User{
		Username: "agilistikmal",
		FullName: "Agil Ghani Istikmal",
		Email:    "agilistikmal3@gmail.com",
		Phone:    "+6281346173829",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)
}

func TestUpdateUserRepository(t *testing.T) {
	userId := "66c1f2915941895aca04faaf"
	user, err := r.Update(&model.UserFilter{
		Id: userId,
	}, &model.User{
		Email: "agil_g@safatanc.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)
}

func TestFindUsersRepository(t *testing.T) {
	users, err := r.FindMany(nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		log.Println(user)
	}
}

func TestFindUserRepository(t *testing.T) {
	filter := &model.UserFilter{
		Username: "agilistikmal",
	}
	user, err := r.FindOne(filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
}

func TestDeleteUserRepository(t *testing.T) {
	filter := &model.UserFilter{
		Email: "agil_g123@safatanc.com",
	}
	user, err := r.Delete(filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)
}
