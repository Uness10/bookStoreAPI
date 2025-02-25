package memory

import (
	"sync"
	"time"

	"bookstore.com/models"
)

type InMemorySalesReportStore struct {
	mu           sync.Mutex
	SalesReports []models.SalesReport
	ordersList   []models.Order
}

var (
	SalesReportInstance *InMemorySalesReportStore
	SalesReportOnce     sync.Once
)

func NewInMemorySalesReportStore() *InMemorySalesReportStore {
	SalesReportOnce.Do(func() {
		SalesReportInstance = &InMemorySalesReportStore{
			SalesReports: make([]models.SalesReport, 0),
		}
	})
	return SalesReportInstance
}

func (s *InMemorySalesReportStore) Create(salesReport models.SalesReport) (models.SalesReport, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.SalesReports = append(s.SalesReports, salesReport)
	return salesReport, nil
}

func (s *InMemorySalesReportStore) Search(query models.SearchCriteria) ([]models.SalesReport, error) {
	var results []models.SalesReport
	if len(query.Filters) == 0 {
		for _, SalesReport := range s.SalesReports {
			results = append(results, SalesReport)
		}
		return results, nil
	}

	for _, SalesReport := range s.SalesReports {
		match := true

		if from, e1 := query.Filters["from"]; e1 {
			if to, e2 := query.Filters["to"]; e2 {
				var err error
				from, err = time.Parse("2021-02-20T08:45:00Z", from.(string))
				to, err = time.Parse("2021-02-20T08:45:00Z", to.(string))
				if err != nil {
					return results, err
				}
			}
		}

		if match {
			results = append(results, SalesReport)
		}
	}

	return results, nil
}
