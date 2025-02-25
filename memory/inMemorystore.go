package memory

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

type InMemoryStore struct {
	BookStore     InMemoryBookStore
	AuthorStore   InMemoryAuthorStore
	CustomerStore InMemoryCustomerStore
	OrderStore    InMemoryOrderStore
	SalesReport   InMemorySalesReportStore
}

var (
	instance *InMemoryStore
	once     sync.Once
	mutex    sync.RWMutex
)

func NewInMemoryStore() (*InMemoryStore, error) {
	var err error
	once.Do(func() {
		instance, err = LoadData()
	})

	if err != nil {
		log.Fatalf("Error loading data: %v", err)
		return nil, err
	}

	// Ensure each store is initialized after loading
	initializeStores(instance)

	return instance, nil
}

func initializeStores(store *InMemoryStore) {
	// Initialize BookStore if it is not initialized
	if store.BookStore.Books == nil {
		store.BookStore = *NewInMemoryBookStore()
	}

	// Initialize AuthorStore if it is not initialized
	if store.AuthorStore.Authors == nil {
		store.AuthorStore = *NewInMemoryAuthorStore()
	}

	// Initialize CustomerStore if it is not initialized
	if store.CustomerStore.Customers == nil {
		store.CustomerStore = *NewInMemoryCustomerStore()
	}

	// Initialize OrderStore if it is not initialized
	if store.OrderStore.Orders == nil {
		store.OrderStore = *NewInMemoryOrderStore()
	}

}
func LoadData() (*InMemoryStore, error) {
	data, err := os.ReadFile("database.json")
	if err != nil {
		// If file doesn't exist, return empty store
		if os.IsNotExist(err) {
			return &InMemoryStore{}, nil
		}
		return nil, err
	}

	store := &InMemoryStore{}
	err = json.Unmarshal(data, &store)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func SaveData(store *InMemoryStore) error {
	mutex.Lock()
	defer mutex.Unlock()

	data, err := json.MarshalIndent(store, "", "    ")
	if err != nil {
		return err
	}

	file, err := os.Create("database.json")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *InMemoryStore) Schedule() {
	go func() {
		for {
			time.Sleep(10 * time.Second)
			err := SaveData(s)
			if err != nil {
				log.Fatal(err)
				return
			}
			log.Println("saving data")
		}

	}()

}
