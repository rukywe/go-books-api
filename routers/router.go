package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rukywe/go-books-api/handlers"
	"github.com/rukywe/go-books-api/middleware"
	"github.com/rukywe/go-books-api/storage"
)

func NewRouter(store *storage.InMemoryStore) *mux.Router {
    router := mux.NewRouter()
    bookHandler := handlers.BookHandler{Store: store}

    router.Use(middleware.Logger)
	router.Use(middleware.RateLimit)
	
    router.Handle("/books", middleware.ValidateBook(http.HandlerFunc(bookHandler.CreateBook))).Methods("POST")
    router.HandleFunc("/books/{id}", bookHandler.GetBook).Methods("GET")
    router.Handle("/books/{id}", middleware.ValidateBook(http.HandlerFunc(bookHandler.UpdateBook))).Methods("PUT")
    router.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods("DELETE")
    router.HandleFunc("/books", bookHandler.ListBooks).Methods("GET")
	
    return router
}
