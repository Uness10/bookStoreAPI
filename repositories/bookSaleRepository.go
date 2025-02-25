package repositories

import (
	"bookstore.com/models"
)

type BookSaleStore interface {
	Create(book models.BookSale) (models.BookSale, error)
	Get(idx int) (models.BookSale, error)
	Delete(idx int) error
	Search(query models.SearchCriteria) ([]models.BookSale, error)
}
