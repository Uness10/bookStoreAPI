package services

import (
	"errors"
	"fmt"

	"bookstore.com/memory"
	"bookstore.com/models"
	"bookstore.com/repositories"
)

type BookService struct {
	bookRepo repositories.BookStore
}

func NewBookService(bookRepo repositories.BookStore) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

// CreateBook adds a new book to the store with validation and context propagation
func (s *BookService) CreateBook(book models.Book) (models.Book, error) {

	_, authorExists := NewAuthorService(memory.NewInMemoryAuthorStore()).GetAuthor(book.Author.ID)
	fmt.Println(book.Author.ID)
	if authorExists != nil {
		return models.Book{}, errors.New("Author not found")
	}
	return s.bookRepo.Create(book)
}

// GetBookByID retrieves a book by its ID, passing context to the repository
func (s *BookService) GetBookByID(id int) (models.Book, error) {
	return s.bookRepo.Get(id)
}

// UpdateBook updates an existing book in the store
func (s *BookService) UpdateBook(book models.Book) (models.Book, error) {

	return s.bookRepo.Update(book)
}

func (s *BookService) DeleteBook(id int) error {
	return s.bookRepo.Delete(id)
}

func (s *BookService) SearchBooks(query models.SearchCriteria) ([]models.Book, error) {
	return s.bookRepo.Search(query)
}
