package main

import (
	"log"
	"net/http"
	"quickstart/api/books_api"
	"quickstart/api/dogs_api"
	products_api "quickstart/api/store_api"
	"quickstart/helper"

	"github.com/gorilla/mux"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	// Books Routes
	r.HandleFunc("/api/books", books_api.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", books_api.GetBook).Methods("GET")
	r.HandleFunc("/api/books", books_api.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", books_api.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", books_api.DeleteBook).Methods("DELETE")
	
	// Dog Routes
	r.HandleFunc("/api/dogs/random", dogs_api.GetRandomDog).Methods("GET")
	
	// Store Routes
	r.HandleFunc("/api/store/products", products_api.GetProducts).Methods("GET")
	r.HandleFunc("/api/store/product", products_api.CreateProduct).Methods("POST")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
