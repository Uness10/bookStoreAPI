package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"bookstore.com/handlers"
	"bookstore.com/memory"
	"bookstore.com/services"
	"github.com/julienschmidt/httprouter"
)

var database *memory.InMemoryStore
var err error

// Initialize database
func init() {
	database, err = memory.NewInMemoryStore()
	if err != nil {
		log.Fatal(err)
	}
}

func DispatcherWrapper(w http.ResponseWriter, r *http.Request, ps httprouter.Params, requestHandler func(http.ResponseWriter, *http.Request, httprouter.Params)) {
	w.Header().Set("Content-Type", "application/json")
	clientContext := r.Context()
	requestContext, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	requestChannel := make(chan bool)

	go func() {
		requestHandler(w, r, ps)
		requestChannel <- true
	}()

	select {
	case <-clientContext.Done():
		log.Println("Connection Lost")
		return
	case <-requestChannel:
		log.Println("Request Done with Success")
	case <-requestContext.Done():
		log.Println("Request Timeout")
	}
}

func main() {
	// Initialize the book service with the in-memory store
	bookHandler := handlers.NewBookHandler(services.NewBookService(&database.BookStore))
	authorHandler := handlers.NewAuthorHandler(services.NewAuthorService(&database.AuthorStore))
	customerHandler := handlers.NewCustomerHandler(services.NewCustomerService(&database.CustomerStore))
	orderHandler := handlers.NewOrderHandler(services.NewOrderService(&database.OrderStore))
	// Set up router
	router := httprouter.New()
	handleBookRequests(router, bookHandler)
	handleAuthorRequests(router, authorHandler)
	handleCustomerRequests(router, customerHandler)
	handleOrderRequests(router, orderHandler)

	//database.Schedule()

	// Start the HTTP server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleBookRequests(router *httprouter.Router, bookHandler *handlers.BookHandler) {
	router.POST("/books", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, bookHandler.CreateBook)
	})
	router.GET("/books/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, bookHandler.GetBookById)
	})
	router.GET("/books", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, bookHandler.GetBooksByCriteria)
	})
	router.PUT("/books/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, bookHandler.UpdateBookById)
	})
	router.DELETE("/books/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, bookHandler.DeleteBookById)
	})

}

func handleAuthorRequests(router *httprouter.Router, authorHandler *handlers.AuthorHandler) {
	router.POST("/authors", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, authorHandler.CreateAuthor)
	})
	router.GET("/authors/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, authorHandler.GetAuthorById)
	})
	router.GET("/authors", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, authorHandler.GetAuthorsByCriteria)
	})
	router.PUT("/authors/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, authorHandler.UpdateAuthorById)
	})
	router.DELETE("/authors/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, authorHandler.DeleteAuthorById)
	})

}

func handleCustomerRequests(router *httprouter.Router, customerHandler *handlers.CustomerHandler) {
	router.POST("/customers", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, customerHandler.CreateCustomer)
	})
	router.GET("/customers/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, customerHandler.GetCustomerById)
	})
	router.GET("/customers", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, customerHandler.GetCustomersByCriteria)
	})
	router.PUT("/customers/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, customerHandler.UpdateCustomerById)
	})
	router.DELETE("/customers/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, customerHandler.DeleteCustomerById)
	})

}
func handleOrderRequests(router *httprouter.Router, orderHandler *handlers.OrderHandler) {
	router.POST("/orders", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, orderHandler.CreateOrder)
	})
	router.GET("/orders/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, orderHandler.GetOrderById)
	})
	router.GET("/orders", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, orderHandler.GetOrdersByCriteria)
	})
	router.PUT("/orders/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, orderHandler.UpdateOrderById)
	})
	router.DELETE("/orders/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		DispatcherWrapper(w, r, ps, orderHandler.DeleteOrderById)
	})

}
