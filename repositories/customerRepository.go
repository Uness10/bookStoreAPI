package repositories

import (
	"bookstore.com/models"
)

type CustomerStore interface {
	Create(Customer models.Customer) (models.Customer, error)
	Get(idx int) (models.Customer, error)
	Update(item models.Customer) (models.Customer, error)
	Delete(idx int) error
	Search(query models.SearchCriteria) ([]models.Customer, error)
}
