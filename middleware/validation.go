package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/rukywe/go-books-api/models"
)

func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func ValidateBook(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var book models.Book
        if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
            writeJSONError(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        if book.Title == "" || book.Author == "" || book.Year == 0 {
            writeJSONError(w, "Missing required book fields", http.StatusBadRequest)
            return
        }

        next.ServeHTTP(w, r)
    })
}
