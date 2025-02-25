package memory

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"bookstore.com/models"
)

type InMemoryAuthorStore struct {
	mu      sync.Mutex
	Authors map[int]models.Author
	nextID  int
}

var (
	authorInstance *InMemoryAuthorStore
	authorOnce     sync.Once
)

func NewInMemoryAuthorStore() *InMemoryAuthorStore {
	authorOnce.Do(func() {
		authorInstance = &InMemoryAuthorStore{
			Authors: make(map[int]models.Author),
			nextID:  1,
		}
	})
	return authorInstance
}

func (s *InMemoryAuthorStore) Create(Author models.Author) (models.Author, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	Author.ID = s.nextID
	s.Authors[s.nextID] = Author
	s.nextID++
	return Author, nil
}

func (s *InMemoryAuthorStore) Get(id int) (models.Author, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	Author, exists := s.Authors[id]
	fmt.Println(s.Authors)
	if !exists {
		return models.Author{}, errors.New("Author not found")
	}
	return Author, nil
}

func (s *InMemoryAuthorStore) Update(Author models.Author) (models.Author, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.Authors[Author.ID]
	if !exists {
		return models.Author{}, errors.New("Author not found")
	}
	s.Authors[Author.ID] = Author
	return Author, nil
}

func (s *InMemoryAuthorStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.Authors[id]
	if !exists {
		return errors.New("Author not found")
	}
	delete(s.Authors, id)
	return nil
}

func (s *InMemoryAuthorStore) Search(query models.SearchCriteria) ([]models.Author, error) {
	var results []models.Author
	if len(query.Filters) == 0 {
		for _, book := range s.Authors {
			results = append(results, book)
		}
		return results, nil
	}
	for _, author := range s.Authors {
		match := true

		if firstName, exists := query.Filters["firstName"]; exists {
			if !strings.Contains(author.FirstName, firstName.(string)) {
				match = false
			}
		}

		if lastName, exists := query.Filters["lastName"]; exists {
			if !strings.Contains(author.LastName, lastName.(string)) {
				match = false
			}
		}

		if name, exists := query.Filters["name"]; exists {
			fullName := author.FirstName + " " + author.LastName
			if !strings.Contains(fullName, name.(string)) {
				match = false
			}
		}

		if match {
			results = append(results, author)
		}
	}

	return results, nil
}
