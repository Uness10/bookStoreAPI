package services

import (
	"errors"

	"bookstore.com/memory"
	"bookstore.com/models"
	"bookstore.com/repositories"
)

type OrderService struct {
	orderRepo repositories.OrderStore
}

func NewOrderService(repo repositories.OrderStore) *OrderService {
	return &OrderService{orderRepo: repo}
}

func (s *OrderService) CreateOrder(order models.Order) (models.Order, error) {
	_, customerExists := NewCustomerService(memory.NewInMemoryCustomerStore()).GetCustomer(order.Customer.ID)
	for _, item := range order.Items {
		_, bookFound := NewOrderItemService(memory.NewInMemoryOrderItemStore()).CreateOrderItem(item)
		if bookFound != nil {
			return models.Order{}, errors.New("Some Books does not exist")
		}
	}
	if customerExists != nil {
		return models.Order{}, errors.New("customer not found")
	}
	return s.orderRepo.Create(order)
}

func (s *OrderService) GetOrder(id int) (models.Order, error) {
	return s.orderRepo.Get(id)
}

func (s *OrderService) UpdateOrder(order models.Order) (models.Order, error) {
	return s.orderRepo.Update(order)
}

func (s *OrderService) DeleteOrder(id int) error {
	return s.orderRepo.Delete(id)
}

func (s *OrderService) SearchOrders(query models.SearchCriteria) ([]models.Order, error) {
	return s.orderRepo.Search(query)
}
