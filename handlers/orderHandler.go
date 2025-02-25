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

// OrderHandler handles order-related HTTP requests.
type OrderHandler struct {
	OrderService *services.OrderService
}

var (
	OrderInstance *OrderHandler
	OrderOnce     sync.Once
)

// NewOrderHandler initializes a singleton instance of OrderHandler.
func NewOrderHandler(OrderService *services.OrderService) *OrderHandler {
	OrderOnce.Do(func() {
		OrderInstance = &OrderHandler{OrderService: OrderService}
	})
	return OrderInstance
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()

	var Order models.Order
	err := json.NewDecoder(r.Body).Decode(&Order)
	if err != nil {
		log.Printf("OrderHandler.Create: invalid input error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdOrder, err := h.OrderService.CreateOrder(Order)
	if err != nil {
		log.Printf("OrderHandler.Create: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdOrder); err != nil {
		log.Printf("OrderHandler.Create: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("OrderHandler.Create: success, duration: %v", time.Since(start))
}

func (h *OrderHandler) GetOrderById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("OrderHandler.GetById: invalid id error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid Order ID", http.StatusBadRequest)
		return
	}

	Order, err := h.OrderService.GetOrder(id)
	if err != nil {
		log.Printf("OrderHandler.GetById: not found error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Order not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Order); err != nil {
		log.Printf("OrderHandler.GetById: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("OrderHandler.GetById: success, duration: %v", time.Since(start))
}

func (h *OrderHandler) GetOrdersByCriteria(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()

	var query = models.SearchCriteria{Filters: make(map[string]interface{})}
	err := json.NewDecoder(r.Body).Decode(&query.Filters)
	if err != nil {
		query.Filters = make(map[string]interface{})
		log.Printf("OrderHandler.Search: invalid criteria error: %v, duration: %v", err, time.Since(start))
	}

	Orders, err := h.OrderService.SearchOrders(query)
	if err != nil {
		log.Printf("OrderHandler.Search: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Orders); err != nil {
		log.Printf("OrderHandler.Search: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("OrderHandler.Search: success, returned %d orders, duration: %v", len(Orders), time.Since(start))
}

func (h *OrderHandler) UpdateOrderById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("OrderHandler.Update: invalid id error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid Order ID", http.StatusBadRequest)
		return
	}

	var Order models.Order
	err = json.NewDecoder(r.Body).Decode(&Order)
	if err != nil {
		log.Printf("OrderHandler.Update: invalid input error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}
	Order.ID = id

	updatedOrder, err := h.OrderService.UpdateOrder(Order)
	if err != nil {
		log.Printf("OrderHandler.Update: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Order not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedOrder); err != nil {
		log.Printf("OrderHandler.Update: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("OrderHandler.Update: success, duration: %v", time.Since(start))
}

func (h *OrderHandler) DeleteOrderById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("OrderHandler.Delete: invalid id error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid Order ID", http.StatusBadRequest)
		return
	}

	err = h.OrderService.DeleteOrder(id)
	if err != nil {
		log.Printf("OrderHandler.Delete: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Order not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Printf("OrderHandler.Delete: success, duration: %v", time.Since(start))
}
