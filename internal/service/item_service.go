package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/repository"
	"github.com/gprestore/gprestore-core/pkg/converter"
	"github.com/gprestore/gprestore-core/pkg/structs"
	"github.com/gprestore/gprestore-core/pkg/variable"
)

type ItemService struct {
	repository *repository.ItemRepository
	validate   *validator.Validate
}

func NewItemService(repository *repository.ItemRepository, validate *validator.Validate) *ItemService {
	return &ItemService{
		repository: repository,
		validate:   validate,
	}
}

func (s *ItemService) Create(input *model.ItemCreate) (*model.Item, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	inputItem, err := converter.StructConverter[model.Item](input)
	if err != nil {
		return nil, err
	}

	item, err := s.repository.Create(inputItem)
	return item, err
}

func (s *ItemService) Update(filter *model.ItemFilter, input *model.ItemUpdate) (*model.Item, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	inputItem, err := converter.StructConverter[model.Item](input)
	if err != nil {
		return nil, err
	}

	item, err := s.repository.Update(filter, inputItem)
	return item, err
}

func (s *ItemService) FindMany(filter *model.ItemFilter) ([]*model.Item, error) {
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

func (s *ItemService) FindOne(filter *model.ItemFilter) (*model.Item, error) {
	isEmpty := structs.IsEmpty(filter)
	if isEmpty {
		return nil, variable.ErrItemFilter
	}

	err := s.validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	item, err := s.repository.FindOne(filter)
	return item, err
}

func (s *ItemService) Delete(filter *model.ItemFilter) (*model.Item, error) {
	err := s.validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	item, err := s.repository.Delete(filter)
	return item, err
}
