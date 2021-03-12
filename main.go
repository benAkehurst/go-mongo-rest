package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"quickstart/helper"
	"quickstart/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = helper.ConnectDB()

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Creates a books array
	var books []models.Book

	// bson.M{}: finds all the books
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// close once cur has finished
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value sor which a single obj can be decoded
		var book models.Book
		// & char returns the memory address of the var
		err := cur.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}

		// add a book to the array of books
		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// encodes the books array into json
	json.NewEncoder(w).Encode(books)
}





func main() {
	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	// r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	// r.HandleFunc("/api/books", createBook).Methods("POST")
	// r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}



// func main() {
// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(ctx)
// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	databases, err := client.ListDatabaseNames(ctx, bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(databases)
// }
