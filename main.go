package main

import (
	"encoding/json"
	"log"
	"net/http"

	"math/rand"
	"strconv"

	"github.com/gorilla/mux"
)

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Book struct {
	ID     int     `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookId, _ := strconv.Atoi(params["id"])
	for _, book := range books {
		if book.ID == bookId {
			json.NewEncoder(w).Encode(book)
			// Return command will escape from function updateBook. No statement after return that will be run
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = rand.Intn(10000000)
	books = append(books, book)
	// Return new book
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookId, _ := strconv.Atoi(params["id"])
	for idx, book := range books {
		if book.ID == bookId {
			// Append semua elemen hingga sebelum index ke "idx", dengan data sesudah index "idx + 1"
			books = append(books[:idx], books[idx+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = rand.Intn(10000000)
			books = append(books, book)
			json.NewEncoder(w).Encode(books)
			// Return command will escape from function updateBook. No statement after return that will be run
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookId, _ := strconv.Atoi(params["id"])
	for idx, book := range books {
		if book.ID == bookId {
			// Append semua elemen hingga sebelum index ke "idx", dengan data sesudah index "idx + 1"
			books = append(books[:idx], books[idx+1:]...)
			// Break the loop and run statement outside the loop
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Mock data - next implement database
	books = append(books, Book{ID: 1, Isbn: "342342342", Title: "Golang Programming", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: 2, Isbn: "342342343", Title: "NodeJs Programming", Author: &Author{FirstName: "Steve", LastName: "Smith"}})
	books = append(books, Book{ID: 3, Isbn: "342342344", Title: "Java Programming", Author: &Author{FirstName: "Steve", LastName: "Smith"}})

	// Route handler
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
