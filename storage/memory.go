package storage

import (
	"sync"

	"github.com/rukywe/go-books-api/models"
)

type InMemoryStore struct {
    books map[string]models.Book
    mu    sync.Mutex
}

func NewInMemoryStore() *InMemoryStore {
    return &InMemoryStore{
        books: make(map[string]models.Book),
    }
}

func (s *InMemoryStore) CreateBook(book models.Book) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.books[book.ID] = book
}

func (s *InMemoryStore) GetBook(id string) (models.Book, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    book, exists := s.books[id]
    return book, exists
}

func (s *InMemoryStore) UpdateBook(book models.Book) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.books[book.ID] = book
}

func (s *InMemoryStore) DeleteBook(id string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    delete(s.books, id)
}

func (s *InMemoryStore) ListBooks() []models.Book {
    s.mu.Lock()
    defer s.mu.Unlock()
    books := make([]models.Book, 0, len(s.books))
    for _, book := range s.books {
        books = append(books, book)
    }
    return books
}
