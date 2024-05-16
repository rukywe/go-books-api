package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
        // Read and print the request body
        body, err := io.ReadAll(r.Body)
        if err != nil {
            writeJSONError(w, "Error reading request body", http.StatusBadRequest)
            return
        }
        log.Printf("Request Body: %s", body)
        
        // Decode the JSON
        var book models.Book
        if err := json.Unmarshal(body, &book); err != nil {
            writeJSONError(w, fmt.Sprintf("Invalid request payload: %v", err), http.StatusBadRequest)
            return
        }

        // Re-assign the request body for further use
        r.Body = io.NopCloser(bytes.NewReader(body))

        // Validate fields
        if book.Title == "" || book.Author == "" || book.Year == 0 {
            writeJSONError(w, "Missing required book fields", http.StatusBadRequest)
            return
        }

        next.ServeHTTP(w, r)
    })
}
