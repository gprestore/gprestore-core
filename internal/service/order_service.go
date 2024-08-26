package service

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/repository"
	"github.com/gprestore/gprestore-core/pkg/converter"
	"github.com/gprestore/gprestore-core/pkg/structs"
	"github.com/gprestore/gprestore-core/pkg/variable"
)

type OrderService struct {
	repository      *repository.OrderRepository
	itemRepository  *repository.ItemRepository
	stockRepository *repository.StockRepository
	paymentService  *PaymentService
	validate        *validator.Validate
}

func NewOrderService(repository *repository.OrderRepository, itemRepository *repository.ItemRepository, stockRepository *repository.StockRepository, paymentService *PaymentService, validate *validator.Validate) *OrderService {
	return &OrderService{
		repository:      repository,
		itemRepository:  itemRepository,
		stockRepository: stockRepository,
		paymentService:  paymentService,
		validate:        validate,
	}
}

func (s *OrderService) Create(input *model.OrderCreate) (*model.Order, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	var newOrderItems []model.OrderItem
	for _, orderItem := range input.Items {
		itemFilter := &model.ItemFilter{
			Id: orderItem.ItemId,
		}

		item, err := s.itemRepository.FindOne(itemFilter)
		if err != nil {
			return nil, err
		}

		if *item.StockCount < orderItem.Quantity {
			return nil, fmt.Errorf("stock for item %v is insufficient", item.Name)
		}

		newOrderItems = append(newOrderItems, model.OrderItem{
			ItemId:   item.Id.Hex(),
			Name:     item.Name,
			Price:    item.Price,
			Quantity: orderItem.Quantity,
		})
	}

	input.Items = newOrderItems

	inputOrder, err := converter.StructConverter[model.Order](input)
	if err != nil {
		return nil, err
	}

	order, err := s.repository.Create(inputOrder)
	if err != nil {
		return nil, err
	}

	paymentChannel, err := s.paymentService.CreatePayment(order)
	if err != nil {
		return nil, err
	}

	orderFilter := &model.OrderFilter{
		Id: order.Id.Hex(),
	}

	orderUpdate := &model.Order{
		PaymentChannel: paymentChannel,
	}

	order, err = s.repository.Update(orderFilter, orderUpdate)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) Update(filter *model.OrderFilter, input *model.OrderUpdate) (*model.Order, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	inputOrder, err := converter.StructConverter[model.Order](input)
	if err != nil {
		return nil, err
	}

	order, err := s.repository.Update(filter, inputOrder)
	return order, err
}

func (s *OrderService) FindMany(filter *model.OrderFilter) ([]*model.Order, error) {
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

func (s *OrderService) FindOne(filter *model.OrderFilter) (*model.Order, error) {
	isEmpty := structs.IsEmpty(filter)
	if isEmpty {
		return nil, variable.ErrOrderFilter
	}

	err := s.validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	order, err := s.repository.FindOne(filter)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) Delete(filter *model.OrderFilter) (*model.Order, error) {
	err := s.validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	order, err := s.repository.Delete(filter)
	return order, err
}
