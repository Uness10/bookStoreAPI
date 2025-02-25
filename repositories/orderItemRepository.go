package repositories

import (
	"bookstore.com/models"
)

type OrderItemStore interface {
	Create(OrderItem models.OrderItem) (models.OrderItem, error)
	Get(idx int) (models.OrderItem, error)
	Update(item models.OrderItem) (models.OrderItem, error)
	Delete(idx int) error
	Search(query models.SearchCriteria) ([]models.OrderItem, error)
}
