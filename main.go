package main

import (
	"log"
	"net/http"
	books_helper "quickstart/books_handler"
	"quickstart/helper"

	"github.com/gorilla/mux"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	// Books Routes
	r.HandleFunc("/api/books", books_helper.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", books_helper.GetBook).Methods("GET")
	r.HandleFunc("/api/books", books_helper.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", books_helper.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", books_helper.DeleteBook).Methods("DELETE")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
