package main

import (
	"log"
	"net/http"

	"bookstore/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// books
	r.HandleFunc("/books", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books", handlers.AddBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	// authors
	r.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
	r.HandleFunc("/authors", handlers.AddAuthor).Methods("POST")

	// categories
	r.HandleFunc("/categories", handlers.GetCategory).Methods("GET")
	r.HandleFunc("/categories", handlers.AddCategory).Methods("POST")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
