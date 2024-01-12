package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"google.golang.org/grpc"
	"GRPC/GRPC" 
	"math/rand"
	"google.golang.org/grpc/codes" 
	"errors"
	"google.golang.org/grpc/status"
)

var (
	mu  sync.Mutex
	db  *gorm.DB
	err error
)

// Book model for REST API
type Book struct {
	gorm.Model
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author  *Author `gorm:"foreignKey:AuthorID" json:"author"`
	AuthorID uint    `json:"-"`
}

// Author model for REST API
type Author struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Book model for gRPC
type GRPCBook struct {
	Id       string       `json:"id"`
	Isbn     string       `json:"isbn"`
	Title    string       `json:"title"`
	Author   *GRPC.Author      `gorm:"foreignKey:AuthorID" json:"author"`
	AuthorID uint         `json:"-"`
}


// REST API handlers

func getBooks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")

	// Perform the database query
	var resultBooks []Book
	if err := db.Preload("Author").Find(&resultBooks).Error; err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resultBooks)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	bookID := params["id"]

	// Perform the database query
	var book Book
	if err := db.Preload("Author").First(&book, "id = ?", bookID).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			json.NewEncoder(w).Encode(&Book{})
		} else {
			log.Fatal(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")

	// Decode values from the request body
	var bookData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&bookData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Extract values from the decoded map
	isbn, ok := bookData["isbn"]
	if !ok {
		http.Error(w, "Missing 'isbn' in request body", http.StatusBadRequest)
		return
	}

	title, ok := bookData["title"]
	if !ok {
		http.Error(w, "Missing 'title' in request body", http.StatusBadRequest)
		return
	}

	authorFirstname, ok := bookData["author_firstname"]
	if !ok {
		http.Error(w, "Missing 'author_firstname' in request body", http.StatusBadRequest)
		return
	}

	authorLastname, ok := bookData["author_lastname"]
	if !ok {
		http.Error(w, "Missing 'author_lastname' in request body", http.StatusBadRequest)
		return
	}

	// Generate a random ID
	id := strconv.Itoa(rand.Intn(10000000))

	// Perform the database insertion for REST API
	newBook := Book{
		ID:     id,
		Isbn:   isbn,
		Title:  title,
		Author: &Author{Firstname: authorFirstname, Lastname: authorLastname},
	}

	if err := db.Create(&newBook).Error; err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&newBook)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// Decode values from the request body
	var updateData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Extract values from the decoded map
	newIsbn, ok := updateData["isbn"]
	if !ok {
		http.Error(w, "Missing 'new_isbn' in request body", http.StatusBadRequest)
		return
	}

	newTitle, ok := updateData["title"]
	if !ok {
		http.Error(w, "Missing 'new_title' in request body", http.StatusBadRequest)
		return
	}

	// Perform the database update for REST API
	if err := db.Model(&Book{}).Where("id = ?", params["id"]).Updates(map[string]interface{}{"isbn": newIsbn, "title": newTitle}).Error; err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Perform the database query to get the updated data
	var resultBooks []Book
	if err := db.Preload("Author").Find(&resultBooks).Error; err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resultBooks)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// Perform the database deletion for REST API
	if err := db.Delete(&Book{}, "id = ?", params["id"]).Error; err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Perform the database query to get the updated data
	var resultBooks []Book
	if err := db.Find(&resultBooks).Error; err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resultBooks)
}

// gRPC handlers

type grpcServer struct {
	GRPC.UnimplementedBookServiceServer
}

func (s *grpcServer) GetBooks(ctx context.Context, req *GRPC.NoRequest) (*GRPC.BookListResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	var grpcBooksPtr []*GRPC.Book

	// Fetch books from the database for gRPC
	if err := db.Find(&grpcBooksPtr).Error; err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return &GRPC.BookListResponse{Books: grpcBooksPtr}, nil
}

func (s *grpcServer) GetBook(ctx context.Context, req *GRPC.BookRequest) (*GRPC.BookResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	// Perform the database query to get the book for gRPC
	var book Book
	if err := db.First(&book, "id = ?", req.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Return nil Book in case the record is not found
			return &GRPC.BookResponse{Book: nil}, nil
		} else {
			log.Println(err)
			return nil, status.Errorf(codes.Internal, "Internal Server Error")
		}
	}

	// Check if ID is an empty string
	if book.ID == "" {
		return &GRPC.BookResponse{Book: nil}, nil
	}

	// Check if Author is nil
	var grpcAuthor *GRPC.Author
	if book.Author != nil {
		grpcAuthor = &GRPC.Author{
			Firstname: book.Author.Firstname,
			Lastname:  book.Author.Lastname,
		}
	}

	// Return the retrieved book directly
	return &GRPC.BookResponse{Book: &GRPC.Book{
		Id:     book.ID,
		Isbn:   book.Isbn,
		Title:  book.Title,
		Author: grpcAuthor,
	}}, nil
}





func (s *grpcServer) CreateBook(ctx context.Context, req *GRPC.Book) (*GRPC.BookResponse, error) {
	mu.Lock()
	defer mu.Unlock()
	// Generate a random ID
	id := strconv.Itoa(rand.Intn(10000000))

	// Create a new book with the provided data
	newBook := Book{
		ID:     id,
		Isbn:   req.Isbn,
		Title:  req.Title,
		Author: &Author{Firstname: req.Author.Firstname, Lastname: req.Author.Lastname},
	}

	// Perform the database insertion for gRPC
	if err := db.Create(&newBook).Error; err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	// Respond with the created book
	return &GRPC.BookResponse{Book: convertToGRPCBook(newBook)}, nil
}

func (s *grpcServer) UpdateBook(ctx context.Context, req *GRPC.BookRequest) (*GRPC.BookListResponse, error) {
    mu.Lock()
    defer mu.Unlock()

    // Find and update the book in the database for gRPC
    if err := db.Model(&Book{}).Where("id = ?", req.Id).Updates(map[string]interface{}{"isbn": "Updated isbn", "title": "Update title"}).Error; err != nil {
        log.Println(err)
        return nil, status.Errorf(codes.Internal, "Internal Server Error")
    }

    // Perform the database query to get the updated book
    var updatedBook Book
    if err := db.First(&updatedBook, "id = ?", req.Id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // If the record is not found, return an empty list
            return &GRPC.BookListResponse{Books: nil}, nil
        }
        log.Println(err)
        return nil, status.Errorf(codes.Internal, "Internal Server Error")
    }

    // Check if updatedBook is not nil before converting
if &updatedBook != nil {
    // Convert the updatedBook to a slice of books for the response
    booksListResponse := &GRPC.BookListResponse{
        Books: []*GRPC.Book{convertToGRPCBook(updatedBook)},
    }

    return booksListResponse, nil
}


    // Return an empty list if updatedBook is nil
    return &GRPC.BookListResponse{Books: nil}, nil
}



// Function to convert Book to GRPC.Book
func convertToGRPCBook(book Book) *GRPC.Book {
	if book.ID == "" {
        return nil
    }
    var grpcAuthor *GRPC.Author
	if book.Author != nil {
		grpcAuthor = &GRPC.Author{
			Firstname: book.Author.Firstname,
			Lastname:  book.Author.Lastname,
		}
	}
	return &GRPC.Book{
		Id:     book.ID,
		Isbn:   book.Isbn,
		Title:  book.Title,
		// Author: &GRPC.Author{Firstname: book.Author.Firstname, Lastname: book.Author.Lastname},
		Author: grpcAuthor,
	}
}


func (s *grpcServer) DeleteBook(ctx context.Context, req *GRPC.BookRequest) (*GRPC.BookListResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	/// Find and delete the book in the database for gRPC
	if err := db.Delete(&Book{}, "id = ?", req.Id).Error; err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	// Respond with the deleted book
	grpcResponse := &GRPC.BookListResponse{
		Books: []*GRPC.Book{
			{
				Id: req.Id,
			},
		},
	}

	return grpcResponse, nil

}

func main() {
	// Set API_TYPE environment variable to "grpc" or "rest" to choose the API type
	apiType := os.Getenv("API_TYPE")
	dbName := os.Getenv("DB_NAME")
	userName := os.Getenv("USER_NAME")
	pw := os.Getenv("PASSWORD")

	// Connect to PostgreSQL database using GORM
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable", userName, pw, dbName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Initialize the database tables
	if err := db.AutoMigrate(&Book{}, &Author{}); err != nil {
		log.Fatal(err)
	}

	// Verify the connection
	if err := db.Exec("SELECT 1").Error; err != nil {
		log.Fatal(err)
	}

	// Initialize gRPC or REST service based on the environment variable
	if apiType == "grpc" {
		// Start gRPC server
		fmt.Println("For Client Side Connection run the command go build client.go and go run client.go ")

		listener, err := net.Listen("tcp", "localhost:8000")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		srv := grpc.NewServer()
		GRPC.RegisterBookServiceServer(srv, &grpcServer{})

		if err := srv.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	} else if apiType == "rest" {
		// Initialize REST service
		router := mux.NewRouter()

		// Sample data for testing
		// books = append(books, Book{ID: "1", Isbn: "12345", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
		// books = append(books, Book{ID: "2", Isbn: "67890", Title: "Book Two", Author: &Author{Firstname: "Jane", Lastname: "Smith"}})

		// REST endpoints
		router.HandleFunc("/api/books", getBooks).Methods("GET")
		router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
		router.HandleFunc("/api/books", createBook).Methods("POST")
		router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
		router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

		// Start the HTTP server
		log.Fatal(http.ListenAndServe("localhost:8080", router))
	} else {
		log.Fatal("Invalid API_TYPE. Set API_TYPE to 'grpc' or 'rest'")
	}
}
