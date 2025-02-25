package services

import (
	"errors"

	"bookstore.com/memory"
	"bookstore.com/models"
	"bookstore.com/repositories"
)

type OrderItemService struct {
	orderItemRepo repositories.OrderItemStore
}

func NewOrderItemService(repo repositories.OrderItemStore) *OrderItemService {
	return &OrderItemService{orderItemRepo: repo}
}

func (s *OrderItemService) CreateOrderItem(orderItem models.OrderItem) (models.OrderItem, error) {
	_, bookExists := NewBookService(memory.NewInMemoryBookStore()).GetBookByID(orderItem.Book.ID)
	if bookExists != nil {
		return models.OrderItem{}, errors.New("book not found")
	}

	return s.orderItemRepo.Create(orderItem)
}

func (s *OrderItemService) GetOrderItem(id int) (models.OrderItem, error) {
	return s.orderItemRepo.Get(id)
}

func (s *OrderItemService) UpdateOrderItem(orderItem models.OrderItem) (models.OrderItem, error) {
	return s.orderItemRepo.Update(orderItem)
}

func (s *OrderItemService) DeleteOrderItem(id int) error {
	return s.orderItemRepo.Delete(id)
}

func (s *OrderItemService) SearchOrderItems(query models.SearchCriteria) ([]models.OrderItem, error) {
	return s.orderItemRepo.Search(query)
}
