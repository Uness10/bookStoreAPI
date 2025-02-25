package repositories

import (
	"bookstore.com/models"
)

type OrderStore interface {
	Create(Order models.Order) (models.Order, error)
	Get(idx int) (models.Order, error)
	Update(item models.Order) (models.Order, error)
	Delete(idx int) error
	Search(query models.SearchCriteria) ([]models.Order, error)
}
