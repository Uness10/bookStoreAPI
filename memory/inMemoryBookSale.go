package memory

import (
	"errors"
	"strings"
	"sync"

	"bookstore.com/models"
)

type InMemoryBookSaleStore struct {
	mu        sync.Mutex
	bookSales map[int]models.BookSale
	nextID    int
}

var (
	bookSaleInstance *InMemoryBookSaleStore
	bookSaleOnce     sync.Once
)

// NewInMemoryBookSaleStore returns a singleton instance of InMemoryBookSaleStore
func NewInMemoryBookSaleStore() *InMemoryBookSaleStore {
	bookSaleOnce.Do(func() {
		bookSaleInstance = &InMemoryBookSaleStore{
			bookSales: make(map[int]models.BookSale),
			nextID:    1,
		}
	})
	return bookSaleInstance
}

// Create adds a new BookSale entry to the store
func (s *InMemoryBookSaleStore) Create(bookSale models.BookSale) (models.BookSale, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	bookSale.ID = s.nextID
	s.bookSales[s.nextID] = bookSale
	s.nextID++
	return bookSale, nil
}

// Get retrieves a BookSale by its ID
func (s *InMemoryBookSaleStore) Get(id int) (models.BookSale, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	bookSale, exists := s.bookSales[id]
	if !exists {
		return models.BookSale{}, errors.New("BookSale not found")
	}
	return bookSale, nil
}

func (s *InMemoryBookSaleStore) Update(bookSale models.BookSale) (models.BookSale, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.bookSales[bookSale.ID]
	if !exists {
		return models.BookSale{}, errors.New("BookSale not found")
	}
	s.bookSales[bookSale.ID] = bookSale
	return bookSale, nil
}

func (s *InMemoryBookSaleStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.bookSales[id]
	if !exists {
		return errors.New("BookSale not found")
	}
	delete(s.bookSales, id)
	return nil
}

func (s *InMemoryBookSaleStore) Search(query models.SearchCriteria) ([]models.BookSale, error) {
	var results []models.BookSale
	if len(query.Filters) == 0 {
		for _, bookSale := range s.bookSales {
			results = append(results, bookSale)
		}
		return results, nil
	}

	for _, bookSale := range s.bookSales {
		match := true

		// Filter by title
		if title, exists := query.Filters["title"]; exists {
			if !strings.Contains(bookSale.Book.Title, title.(string)) {
				match = false
			}
		}

		// Filter by author (assuming you want the author's first name)
		if author, exists := query.Filters["author"]; exists {
			if !strings.Contains(bookSale.Book.Author.FirstName, author.(string)) {
				match = false
			}
		}

		// Filter by genre
		if genre, exists := query.Filters["genre"]; exists {
			genreMatch := false
			for _, g := range bookSale.Book.Genres {
				if strings.Contains(g, genre.(string)) {
					genreMatch = true
					break
				}
			}
			if !genreMatch {
				match = false
			}
		}

		// Filter by quantity
		if quantity, exists := query.Filters["quantity"]; exists {
			if bookSale.Quantity != quantity {
				match = false
			}
		}

		// If the book sale matches all filters, add to results
		if match {
			results = append(results, bookSale)
		}
	}

	return results, nil
}
