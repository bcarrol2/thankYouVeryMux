package main

import (
	"encoding/json"
	// "log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName string	`json:"last_name"`
}

// init books var as a slice Book struct
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // gets any params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "425277", Title: "It", Author: &Author{FirstName: "Stephen", LastName: "King"}})
	books = append(books, Book{ID: "2", Isbn: "425249", Title: "The Long Walk", Author: &Author{FirstName: "Stephen", LastName: "King"}})

	router.HandleFunc("/api/v1/books", getBooks).Methods("GET")
	router.HandleFunc("/api/v1/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/v1/books", createBook).Methods("POST")
	router.HandleFunc("/api/v1/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/v1/books/{id}", deleteBook).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}