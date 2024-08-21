package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/repository"
	"github.com/gprestore/gprestore-core/pkg/converter"
	"github.com/gprestore/gprestore-core/pkg/structs"
	"github.com/gprestore/gprestore-core/pkg/variable"
)

type StoreService struct {
	repository *repository.StoreRepository
	validate   *validator.Validate
}

func NewStoreService(repository *repository.StoreRepository, validate *validator.Validate) *StoreService {
	return &StoreService{
		repository: repository,
		validate:   validate,
	}
}

func (s *StoreService) Create(input *model.StoreCreate) (*model.Store, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	inputStore, err := converter.StructConverter[model.Store](input)
	if err != nil {
		return nil, err
	}

	store, err := s.repository.Create(inputStore)
	return store, err
}

func (s *StoreService) Update(filter *model.StoreFilter, input *model.StoreUpdate) (*model.Store, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	inputStore, err := converter.StructConverter[model.Store](input)
	if err != nil {
		return nil, err
	}

	store, err := s.repository.Update(filter, inputStore)
	return store, err
}

func (s *StoreService) FindMany(filter *model.StoreFilter) ([]*model.Store, error) {
	err := s.validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	isEmpty := structs.IsEmpty(filter)
	if isEmpty {
		filter = nil
	}

	stores, err := s.repository.FindMany(filter)
	return stores, err

}

func (s *StoreService) FindOne(filter *model.StoreFilter) (*model.Store, error) {
	isEmpty := structs.IsEmpty(filter)
	if isEmpty {
		return nil, variable.ErrFilter
	}

	err := s.validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	store, err := s.repository.FindOne(filter)
	return store, err
}

func (s *StoreService) Delete(filter *model.StoreFilter) (*model.Store, error) {
	err := s.validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	store, err := s.repository.Delete(filter)
	return store, err
}
