package repositories

import (
	"bookstore.com/models"
)

type BookStore interface {
	Create(book models.Book) (models.Book, error)

	Get(idx int) (models.Book, error)

	Update(item models.Book) (models.Book, error)

	Delete(idx int) error

	Search(query models.SearchCriteria) ([]models.Book, error)
}
