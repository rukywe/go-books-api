package main

import (
	"log"
	"net/http"

	"github.com/rukywe/go-books-api/routers"
	"github.com/rukywe/go-books-api/storage"
)

func main() {
    store := storage.NewInMemoryStore()
    router := routers.NewRouter(store)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("could not start server: %v\n", err)
    }
}
