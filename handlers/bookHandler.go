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

type BookHandler struct {
	bookService *services.BookService
}

var (
	bookInstance *BookHandler
	bookOnce     sync.Once
)

func NewBookHandler(bookService *services.BookService) *BookHandler {
	bookOnce.Do(func() {
		bookInstance = &BookHandler{bookService: bookService}
	})
	return bookInstance
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Printf("BookHandler.Create: invalid input error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdBook, err := h.bookService.CreateBook(book)
	if err != nil {
		log.Printf("BookHandler.Create: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdBook); err != nil {
		log.Printf("BookHandler.Create: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("BookHandler.Create: success, duration: %v", time.Since(start))
}

func (h *BookHandler) GetBookById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("BookHandler.GetById: invalid id error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := h.bookService.GetBookByID(id)
	if err != nil {
		log.Printf("BookHandler.GetById: not found error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Book not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		log.Printf("BookHandler.GetById: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("BookHandler.GetById: success, duration: %v", time.Since(start))
}

func (h *BookHandler) GetBooksByCriteria(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()

	var query = models.SearchCriteria{Filters: make(map[string]interface{})}
	if err := json.NewDecoder(r.Body).Decode(&query.Filters); err != nil {
		query.Filters = make(map[string]interface{})
		log.Printf("BookHandler.Search: invalid criteria error: %v, duration: %v", err, time.Since(start))
	}

	books, err := h.bookService.SearchBooks(query)
	if err != nil {
		log.Printf("BookHandler.Search: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		log.Printf("BookHandler.Search: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("BookHandler.Search: success, returned %d books, duration: %v", len(books), time.Since(start))
}

func (h *BookHandler) UpdateBookById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("BookHandler.Update: invalid id error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	if err = json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Printf("BookHandler.Update: invalid input error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}
	book.ID = id

	updatedBook, err := h.bookService.UpdateBook(book)
	if err != nil {
		log.Printf("BookHandler.Update: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Book not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedBook); err != nil {
		log.Printf("BookHandler.Update: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("BookHandler.Update: success, duration: %v", time.Since(start))
}

func (h *BookHandler) DeleteBookById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("BookHandler.Delete: invalid id error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	if err = h.bookService.DeleteBook(id); err != nil {
		log.Printf("BookHandler.Delete: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Book not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Printf("BookHandler.Delete: success, duration: %v", time.Since(start))
}
