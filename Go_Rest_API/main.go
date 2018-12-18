package main

import (
	"encoding/json"
	"log"
	"net/http"

	//"math/rand"
	//"strconv"
	"github.com/gorilla/mux"
)

//Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
	Year   int     `json:"year"`
}

//Author Struct used in the Book Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	age       int    `json:"age"`
}

//Init books var as a slice (array with variable size) book struct
var books []Book

//Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

//Get a single book
func getBook(w http.ResponseWriter, r *http.Request) {

}

//Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {

}

//Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//Delete a books
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//init router
	r := mux.NewRouter()

	//Mock data @todo - Database implementation
	books = append(books, Book{ID: "1", Isbn: "23421364", Title: "Lord of The Rings",
		Author: &Author{Firstname: "J.R.", Lastname: "Tolkien", age: 56}, Year: 1976})
	books = append(books, Book{ID: "2", Isbn: "67856433", Title: "Blade Runner",
		Author: &Author{Firstname: "Philip K.", Lastname: "Dick", age: 65}, Year: 1987})

	//Route Handlers /Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//to run the server
	log.Fatal(http.ListenAndServe(":8000", r))

}
