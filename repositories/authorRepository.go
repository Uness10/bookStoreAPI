package repositories

import (
	"bookstore.com/models"
)

type AuthorStore interface {
	Create(Author models.Author) (models.Author, error)
	Get(idx int) (models.Author, error)
	Update(item models.Author) (models.Author, error)
	Delete(idx int) error
	Search(query models.SearchCriteria) ([]models.Author, error)
}
