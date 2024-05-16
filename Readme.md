# Go Books API

A simple RESTful API for managing a collection of books, built with Go. This project demonstrates basic CRUD operations using an in-memory store and follows Go best practices and project structure.

## Project Structure

1. **Clone the repository:**

```bash
 git clone https://github.com/rukywe/go-books-api.git
 cd go-books-api
```

2. **Initialize the Go module:**

```bash
go mod tidy

```

3. **Initialize the Go module:**

This project uses gorilla/mux for routing and google/uuid for generating unique IDs. You can install them by running:

```bash
go get github.com/gorilla/mux
go get github.com/google/uuid

```

## Running the API

To start the server, run:

```bash
go run main.go
```

The server will start on http://localhost:8080.
