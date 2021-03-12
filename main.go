package main

import (
	"log"
	"net/http"
	books_handler "quickstart/books_handler"
	"quickstart/dogs_handler"
	"quickstart/helper"

	"github.com/gorilla/mux"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	// Books Routes
	r.HandleFunc("/api/books", books_handler.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", books_handler.GetBook).Methods("GET")
	r.HandleFunc("/api/books", books_handler.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", books_handler.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", books_handler.DeleteBook).Methods("DELETE")

	// Dog routes
	r.HandleFunc("/api/dogs/random", dogs_handler.GetRandomDog).Methods("GET")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
