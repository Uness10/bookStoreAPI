package services

import (
	"bookstore.com/models"
	"bookstore.com/repositories"
)

type CustomerService struct {
	customerRepo repositories.CustomerStore
}

func NewCustomerService(repo repositories.CustomerStore) *CustomerService {
	return &CustomerService{customerRepo: repo}
}

func (s *CustomerService) CreateCustomer(customer models.Customer) (models.Customer, error) {
	return s.customerRepo.Create(customer)
}

func (s *CustomerService) GetCustomer(id int) (models.Customer, error) {
	return s.customerRepo.Get(id)
}

func (s *CustomerService) UpdateCustomer(customer models.Customer) (models.Customer, error) {
	return s.customerRepo.Update(customer)
}

func (s *CustomerService) DeleteCustomer(id int) error {
	return s.customerRepo.Delete(id)
}

func (s *CustomerService) SearchCustomers(query models.SearchCriteria) ([]models.Customer, error) {
	return s.customerRepo.Search(query)
}
