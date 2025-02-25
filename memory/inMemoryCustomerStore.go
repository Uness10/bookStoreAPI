package memory

import (
	"errors"
	"fmt"
	"sync"

	"bookstore.com/models"
)

type InMemoryCustomerStore struct {
	mu        sync.Mutex
	Customers map[int]models.Customer
	nextID    int
}

var (
	customerStoreInstance *InMemoryCustomerStore
	customerStoreOnce     sync.Once
)

// NewInMemoryCustomerStore returns the singleton instance of InMemoryCustomerStore
func NewInMemoryCustomerStore() *InMemoryCustomerStore {
	// Initialize the singleton instance only once
	customerStoreOnce.Do(func() {
		customerStoreInstance = &InMemoryCustomerStore{
			Customers: make(map[int]models.Customer),
			nextID:    1,
		}
	})
	return customerStoreInstance
}

// Create adds a new customer to the store
func (s *InMemoryCustomerStore) Create(Customer models.Customer) (models.Customer, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	Customer.ID = s.nextID
	s.Customers[s.nextID] = Customer
	s.nextID++
	return Customer, nil
}

// Get retrieves a customer by ID
func (s *InMemoryCustomerStore) Get(id int) (models.Customer, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Println(s.Customers)
	Customer, exists := s.Customers[id]
	if !exists {
		return models.Customer{}, errors.New("Customer not found")
	}
	return Customer, nil
}

// Update modifies an existing customer in the store
func (s *InMemoryCustomerStore) Update(Customer models.Customer) (models.Customer, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.Customers[Customer.ID]
	if !exists {
		return models.Customer{}, errors.New("Customer not found")
	}
	s.Customers[Customer.ID] = Customer
	return Customer, nil
}

// Delete removes a customer by ID
func (s *InMemoryCustomerStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.Customers[id]
	if !exists {
		return errors.New("Customer not found")
	}
	delete(s.Customers, id)
	return nil
}

// Search filters customers based on the search criteria
func (s *InMemoryCustomerStore) Search(query models.SearchCriteria) ([]models.Customer, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var results []models.Customer
	for _, Customer := range s.Customers {
		results = append(results, Customer)
	}
	return results, nil
}
