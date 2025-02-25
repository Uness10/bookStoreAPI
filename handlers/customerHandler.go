package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"bookstore.com/models"
	"bookstore.com/services"
	"github.com/julienschmidt/httprouter"
)

// CustomerHandler handles customer-related HTTP requests.
type CustomerHandler struct {
	CustomerService *services.CustomerService
}

var (
	CustomerInstance *CustomerHandler
	CustomerOnce     sync.Once
)

// NewCustomerHandler initializes a singleton instance of CustomerHandler.
func NewCustomerHandler(CustomerService *services.CustomerService) *CustomerHandler {
	CustomerOnce.Do(func() {
		CustomerInstance = &CustomerHandler{CustomerService: CustomerService}
	})
	return CustomerInstance
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()

	var Customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&Customer)
	if err != nil {
		log.Printf("CustomerHandler.Create: invalid input error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdCustomer, err := h.CustomerService.CreateCustomer(Customer)
	if err != nil {
		log.Printf("CustomerHandler.Create: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdCustomer); err != nil {
		log.Printf("CustomerHandler.Create: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("CustomerHandler.Create: success, duration: %v", time.Since(start))
}

func (h *CustomerHandler) GetCustomerById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("CustomerHandler.GetById: invalid id error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid Customer ID", http.StatusBadRequest)
		return
	}

	Customer, err := h.CustomerService.GetCustomer(id)
	if err != nil {
		log.Printf("CustomerHandler.GetById: not found error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Customer not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Customer); err != nil {
		log.Printf("CustomerHandler.GetById: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("CustomerHandler.GetById: success, duration: %v", time.Since(start))
}

func (h *CustomerHandler) GetCustomersByCriteria(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()

	var query = models.SearchCriteria{Filters: make(map[string]interface{})}
	err := json.NewDecoder(r.Body).Decode(&query.Filters)
	if err != nil {
		query.Filters = make(map[string]interface{})
		log.Printf("CustomerHandler.Search: invalid criteria error: %v, duration: %v", err, time.Since(start))
	}

	Customers, err := h.CustomerService.SearchCustomers(query)
	if err != nil {
		log.Printf("CustomerHandler.Search: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Customers); err != nil {
		log.Printf("CustomerHandler.Search: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("CustomerHandler.Search: success, returned %d customers, duration: %v", len(Customers), time.Since(start))
}

func (h *CustomerHandler) UpdateCustomerById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("CustomerHandler.Update: invalid id error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid Customer ID", http.StatusBadRequest)
		return
	}

	var Customer models.Customer
	err = json.NewDecoder(r.Body).Decode(&Customer)
	if err != nil {
		log.Printf("CustomerHandler.Update: invalid input error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}
	Customer.ID = id

	updatedCustomer, err := h.CustomerService.UpdateCustomer(Customer)
	if err != nil {
		log.Printf("CustomerHandler.Update: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Customer not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedCustomer); err != nil {
		log.Printf("CustomerHandler.Update: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("CustomerHandler.Update: success, duration: %v", time.Since(start))
}

func (h *CustomerHandler) DeleteCustomerById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("CustomerHandler.Delete: invalid id error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid Customer ID", http.StatusBadRequest)
		return
	}

	err = h.CustomerService.DeleteCustomer(id)
	if err != nil {
		log.Printf("CustomerHandler.Delete: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Customer not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Printf("CustomerHandler.Delete: success, duration: %v", time.Since(start))
}
