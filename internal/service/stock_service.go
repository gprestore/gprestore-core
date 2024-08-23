package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/repository"
	"github.com/gprestore/gprestore-core/pkg/converter"
	"github.com/gprestore/gprestore-core/pkg/structs"
	"github.com/gprestore/gprestore-core/pkg/variable"
)

type StockService struct {
	repository *repository.StockRepository
	validate   *validator.Validate
}

func NewStockService(repository *repository.StockRepository, validate *validator.Validate) *StockService {
	return &StockService{
		repository: repository,
		validate:   validate,
	}
}

func (s *StockService) Create(input *model.StockCreate) (*model.Stock, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	inputStock, err := converter.StructConverter[model.Stock](input)
	if err != nil {
		return nil, err
	}

	item, err := s.repository.Create(inputStock)
	return item, err
}

func (s *StockService) Update(filter *model.StockFilter, input *model.StockUpdate) (*model.Stock, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	inputStock, err := converter.StructConverter[model.Stock](input)
	if err != nil {
		return nil, err
	}

	item, err := s.repository.Update(filter, inputStock)
	return item, err
}

func (s *StockService) FindMany(filter *model.StockFilter) ([]*model.Stock, error) {
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

func (s *StockService) FindOne(filter *model.StockFilter) (*model.Stock, error) {
	isEmpty := structs.IsEmpty(filter)
	if isEmpty {
		return nil, variable.ErrStockFilter
	}

	err := s.validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	item, err := s.repository.FindOne(filter)
	return item, err
}

func (s *StockService) Delete(filter *model.StockFilter) (*model.Stock, error) {
	err := s.validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	item, err := s.repository.Delete(filter)
	return item, err
}
