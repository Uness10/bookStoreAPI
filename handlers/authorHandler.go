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

type AuthorHandler struct {
	AuthorService *services.AuthorService
}

var (
	AuthorInstance *AuthorHandler
	AuthorOnce     sync.Once
)

func NewAuthorHandler(AuthorService *services.AuthorService) *AuthorHandler {
	AuthorOnce.Do(func() {
		AuthorInstance = &AuthorHandler{AuthorService: AuthorService}
	})
	return AuthorInstance
}

func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()

	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		log.Printf("AuthorHandler.Create: invalid input error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdAuthor, err := h.AuthorService.CreateAuthor(author)
	if err != nil {
		log.Printf("AuthorHandler.Create: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdAuthor); err != nil {
		log.Printf("AuthorHandler.Create: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("AuthorHandler.Create: success, duration: %v", time.Since(start))
}

func (h *AuthorHandler) GetAuthorById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("AuthorHandler.GetById: invalid id error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid Author ID", http.StatusBadRequest)
		return
	}

	author, err := h.AuthorService.GetAuthor(id)
	if err != nil {
		log.Printf("AuthorHandler.GetById: not found error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Author not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(author); err != nil {
		log.Printf("AuthorHandler.GetById: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("AuthorHandler.GetById: success, duration: %v", time.Since(start))
}

func (h *AuthorHandler) GetAuthorsByCriteria(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()

	var query = models.SearchCriteria{Filters: make(map[string]interface{})}
	if err := json.NewDecoder(r.Body).Decode(&query.Filters); err != nil {
		query.Filters = make(map[string]interface{})
		log.Printf("AuthorHandler.Search: invalid criteria error: %v, duration: %v", err, time.Since(start))
	}

	authors, err := h.AuthorService.SearchAuthors(query)
	if err != nil {
		log.Printf("AuthorHandler.Search: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(authors); err != nil {
		log.Printf("AuthorHandler.Search: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("AuthorHandler.Search: success, returned %d authors, duration: %v", len(authors), time.Since(start))
}

func (h *AuthorHandler) UpdateAuthorById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("AuthorHandler.Update: invalid id error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid Author ID", http.StatusBadRequest)
		return
	}

	var author models.Author
	if err = json.NewDecoder(r.Body).Decode(&author); err != nil {
		log.Printf("AuthorHandler.Update: invalid input error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}
	author.ID = id

	updatedAuthor, err := h.AuthorService.UpdateAuthor(author)
	if err != nil {
		log.Printf("AuthorHandler.Update: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Author not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedAuthor); err != nil {
		log.Printf("AuthorHandler.Update: encoding error: %v, duration: %v", err, time.Since(start))
		return
	}

	log.Printf("AuthorHandler.Update: success, duration: %v", time.Since(start))
}

func (h *AuthorHandler) DeleteAuthorById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("AuthorHandler.Delete: invalid id error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Invalid Author ID", http.StatusBadRequest)
		return
	}

	if err = h.AuthorService.DeleteAuthor(id); err != nil {
		log.Printf("AuthorHandler.Delete: service error: %v, duration: %v", err, time.Since(start))
		http.Error(w, "Author not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Printf("AuthorHandler.Delete: success, duration: %v", time.Since(start))
}
