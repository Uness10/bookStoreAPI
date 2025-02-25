package services

import (
	"bookstore.com/models"
	"bookstore.com/repositories"
)

type AuthorService struct {
	authorRepo repositories.AuthorStore
}

func NewAuthorService(repo repositories.AuthorStore) *AuthorService {
	return &AuthorService{authorRepo: repo}
}

func (s *AuthorService) CreateAuthor(author models.Author) (models.Author, error) {
	return s.authorRepo.Create(author)
}

func (s *AuthorService) GetAuthor(id int) (models.Author, error) {
	return s.authorRepo.Get(id)
}

func (s *AuthorService) UpdateAuthor(author models.Author) (models.Author, error) {
	return s.authorRepo.Update(author)
}

func (s *AuthorService) DeleteAuthor(id int) error {
	return s.authorRepo.Delete(id)
}

func (s *AuthorService) SearchAuthors(query models.SearchCriteria) ([]models.Author, error) {
	return s.authorRepo.Search(query)
}
