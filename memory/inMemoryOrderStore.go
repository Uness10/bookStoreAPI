package memory

import (
	"errors"
	"sync"

	"bookstore.com/models"
)

type InMemoryOrderStore struct {
	mu     sync.Mutex
	Orders map[int]models.Order
	nextID int
}

var (
	orderStoreInstance *InMemoryOrderStore
	orderStoreOnce     sync.Once
)

// NewInMemoryOrderStore returns the singleton instance of InMemoryOrderStore
func NewInMemoryOrderStore() *InMemoryOrderStore {
	// Initialize the singleton instance only once
	orderStoreOnce.Do(func() {
		orderStoreInstance = &InMemoryOrderStore{
			Orders: make(map[int]models.Order),
			nextID: 1,
		}
	})
	return orderStoreInstance
}

// Create adds a new order to the store
func (s *InMemoryOrderStore) Create(Order models.Order) (models.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	Order.ID = s.nextID
	s.Orders[s.nextID] = Order

	s.nextID++
	NewInMemorySalesReportStore().ordersList = append(NewInMemorySalesReportStore().ordersList, Order)

	return Order, nil
}

// Get retrieves an order by ID
func (s *InMemoryOrderStore) Get(id int) (models.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	Order, exists := s.Orders[id]
	if !exists {
		return models.Order{}, errors.New("Order not found")
	}
	return Order, nil
}

// Update modifies an existing order in the store
func (s *InMemoryOrderStore) Update(Order models.Order) (models.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.Orders[Order.ID]
	if !exists {
		return models.Order{}, errors.New("Order not found")
	}
	s.Orders[Order.ID] = Order
	return Order, nil
}

// Delete removes an order by ID
func (s *InMemoryOrderStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.Orders[id]
	if !exists {
		return errors.New("Order not found")
	}
	delete(s.Orders, id)
	return nil
}

// Search filters orders based on the search criteria
func (s *InMemoryOrderStore) Search(query models.SearchCriteria) ([]models.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var results []models.Order
	for _, Order := range s.Orders {
		results = append(results, Order)
	}
	return results, nil
}
