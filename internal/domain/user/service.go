package user

import (
	"github.com/gprestore/gprestore-core/internal/model"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) Create(input *model.User) (*model.User, error) {
	return nil, nil
}

func (s *Service) Update(id string, input *model.User) (*model.User, error) {
	return nil, nil
}

func (s *Service) FindMany() ([]*model.User, error) {
	return nil, nil
}

func (s *Service) FindOne(id string) (*model.User, error) {
	return nil, nil
}

func (s *Service) Delete(id string) (*model.User, error) {
	return nil, nil
}
