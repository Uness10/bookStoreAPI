package services

import (
	"bookstore.com/models"
	"bookstore.com/repositories"
)

type BookSaleService struct {
	BookSaleRepo repositories.BookSaleStore
}

func NewBookSaleService(repo repositories.BookSaleStore) *BookSaleService {
	return &BookSaleService{BookSaleRepo: repo}
}

func (s *BookSaleService) CreateBookSale(BookSale models.BookSale) (models.BookSale, error) {
	return s.BookSaleRepo.Create(BookSale)
}

func (s *BookSaleService) GetBookSale(id int) (models.BookSale, error) {
	return s.BookSaleRepo.Get(id)
}

func (s *BookSaleService) DeleteBookSale(id int) error {
	return s.BookSaleRepo.Delete(id)
}

func (s *BookSaleService) SearchBookSales(query models.SearchCriteria) ([]models.BookSale, error) {
	return s.BookSaleRepo.Search(query)
}
