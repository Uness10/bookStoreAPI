package memory

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"bookstore.com/models"
)

type InMemoryBookStore struct {
	mu     sync.Mutex
	Books  map[int]models.Book
	nextID int
}

var (
	bookStoreInstance *InMemoryBookStore
	bookStoreOnce     sync.Once
)

// NewInMemoryBookStore returns the singleton instance of InMemoryBookStore
func NewInMemoryBookStore() *InMemoryBookStore {
	// Initialize the singleton instance only once
	bookStoreOnce.Do(func() {
		bookStoreInstance = &InMemoryBookStore{
			Books:  make(map[int]models.Book),
			nextID: 1,
		}
	})
	return bookStoreInstance
}

// Create adds a new book to the store
func (s *InMemoryBookStore) Create(book models.Book) (models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	book.ID = s.nextID
	s.Books[s.nextID] = book
	fmt.Println(s.Books)
	s.nextID++
	return book, nil
}

// Get retrieves a book by ID
func (s *InMemoryBookStore) Get(id int) (models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	book, exists := s.Books[id]
	if !exists {
		return models.Book{}, errors.New("book not found")
	}
	return book, nil
}

// Update modifies an existing book in the store
func (s *InMemoryBookStore) Update(book models.Book) (models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.Books[book.ID]
	if !exists {
		return models.Book{}, errors.New("book not found")
	}
	s.Books[book.ID] = book
	return book, nil
}

// Delete removes a book by ID
func (s *InMemoryBookStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.Books[id]
	if !exists {
		return errors.New("book not found")
	}
	delete(s.Books, id)
	return nil
}

// Search filters books based on the search criteria
func (s *InMemoryBookStore) Search(query models.SearchCriteria) ([]models.Book, error) {
	var results []models.Book
	if len(query.Filters) == 0 {
		for _, book := range s.Books {
			results = append(results, book)
		}
		return results, nil
	}
	for _, book := range s.Books {
		match := true

		if title, exists := query.Filters["title"]; exists {
			if !strings.Contains(book.Title, title.(string)) {
				match = false
			}
		}

		if author, exists := query.Filters["author"]; exists {
			if !strings.Contains(book.Author.FirstName, author.(string)) {
				match = false
			}
		}

		if genre, exists := query.Filters["genre"]; exists {
			genreMatch := false
			for _, g := range book.Genres {
				if strings.Contains(g, genre.(string)) {
					genreMatch = true
					break
				}
			}
			if !genreMatch {
				match = false
			}
		}

		if price, exists := query.Filters["price"]; exists {
			if book.Price != price.(float64) {
				match = false
			}
		}

		if match {
			results = append(results, book)
		}
	}

	return results, nil
}
