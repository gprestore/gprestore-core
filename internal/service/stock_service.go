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
	repository     *repository.StockRepository
	itemRepository *repository.ItemRepository
	validate       *validator.Validate
}

func NewStockService(repository *repository.StockRepository, itemRepository *repository.ItemRepository, validate *validator.Validate) *StockService {
	return &StockService{
		repository:     repository,
		itemRepository: itemRepository,
		validate:       validate,
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

	stock, err := s.repository.Create(inputStock)
	return stock, err
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

	stock, err := s.repository.Update(filter, inputStock)
	if err != nil {
		return nil, err
	}

	itemFilter := &model.ItemFilter{
		Id: stock.ItemId,
	}

	itemUpdate := &model.ItemUpdate{
		StockCount: stock.Count,
	}

	itemInput, err := converter.StructConverter[model.Item](itemUpdate)
	if err != nil {
		return nil, err
	}

	_, err = s.itemRepository.Update(itemFilter, itemInput)
	if err != nil {
		return nil, err
	}

	return stock, nil
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

	stock, err := s.repository.FindOne(filter)
	return stock, err
}

func (s *StockService) Delete(filter *model.StockFilter) (*model.Stock, error) {
	err := s.validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	stock, err := s.repository.Delete(filter)
	return stock, err
}
