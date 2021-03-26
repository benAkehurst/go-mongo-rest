package books_api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"quickstart/helper"
	"quickstart/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


var collection = helper.ConnectDB().Collection("books")

func GetBooks(w http.ResponseWriter, r *http.Request) {
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

func GetBook(w http.ResponseWriter, r *http.Request) {
	// set the header
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	// get route params with mux
	var params = mux.Vars(r)

	// string to primitive.ObjectId
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// creates a filter
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book

	// decode body of request
	_ = json.NewDecoder(r.Body).Decode(&book)

	// insert book model into collection
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var book models.Book

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&book)

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"isbn", book.Isbn},
			{"title", book.Title},
			{"author", bson.D{
				{"firstname", book.Author.FirstName},
				{"lastname", book.Author.LastName},
			}},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	book.ID = id

	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Set headers
	w.Header().Set("Content-Type", "application/json")

	// get route params
	var params = mux.Vars(r)

	// convert param to objectid
	id, err := primitive.ObjectIDFromHex(params["id"])

	// prepare filter
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}
