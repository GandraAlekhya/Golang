package main

import (
	//"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	//"strconv"
	"encoding/json"
	//"math/rand"
	"github.com/gorilla/mux"
)

var books []Book

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

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

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {

		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break

		}

		json.NewEncoder(w).Encode(books)

	}

}

func main() {

	r := mux.NewRouter()

	books = append(books, Book{"1", "abc", &Author{"alekhya", "gandra"}})
	books = append(books, Book{"2", "abcd", &Author{"alekhya1", "gandra1"}})
	books = append(books, Book{"3", "abce", &Author{"alekhya2", "gandra2"}})
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	//	r.HandleFunc("/api/books/{id}", createBook).Methods("POST")
	//r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
