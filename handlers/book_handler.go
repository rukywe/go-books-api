package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rukywe/go-books-api/models"
	"github.com/rukywe/go-books-api/storage"
)

type BookHandler struct {
    Store *storage.InMemoryStore
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    book.ID = uuid.New().String()
    h.Store.CreateBook(book)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Book created", "id": book.ID})
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    book, exists := h.Store.GetBook(id)
    if !exists {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    h.Store.UpdateBook(book)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Book updated"})
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    h.Store.DeleteBook(id)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted"})
}

func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
    books := h.Store.ListBooks()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}
