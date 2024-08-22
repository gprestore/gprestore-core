package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/repository"
	"github.com/gprestore/gprestore-core/pkg/converter"
	"github.com/gprestore/gprestore-core/pkg/structs"
	"github.com/gprestore/gprestore-core/pkg/variable"
)

type UserService struct {
	repository *repository.UserRepository
	validate   *validator.Validate
}

func NewUserService(repository *repository.UserRepository, validate *validator.Validate) *UserService {
	return &UserService{
		repository: repository,
		validate:   validate,
	}
}

func (s *UserService) Create(input *model.UserCreate) (*model.User, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	inputUser, err := converter.StructConverter[model.User](input)
	if err != nil {
		return nil, err
	}

	user, err := s.repository.Create(inputUser)
	return user, err
}

func (s *UserService) Update(filter *model.UserFilter, input *model.UserUpdate) (*model.User, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	inputUser, err := converter.StructConverter[model.User](input)
	if err != nil {
		return nil, err
	}

	user, err := s.repository.Update(filter, inputUser)
	return user, err
}

func (s *UserService) FindMany(filter *model.UserFilter) ([]*model.User, error) {
	err := s.validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	isEmpty := structs.IsEmpty(filter)
	if isEmpty {
		filter = nil
	}

	users, err := s.repository.FindMany(filter)
	return users, err

}

func (s *UserService) FindOne(filter *model.UserFilter) (*model.User, error) {
	isEmpty := structs.IsEmpty(filter)
	if isEmpty {
		return nil, variable.ErrUserFilter
	}

	err := s.validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	user, err := s.repository.FindOne(filter)
	return user, err
}

func (s *UserService) Delete(filter *model.UserFilter) (*model.User, error) {
	err := s.validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	user, err := s.repository.Delete(filter)
	return user, err
}
