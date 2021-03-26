package products_api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"quickstart/helper"
	"quickstart/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = helper.ConnectDB().Collection("products")

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Creates a product array
	var products []models.Product

	//bson.M{}: finds all the books
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// closes connection after searching DB
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var product models.Product
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product

	// decodes the body of the request
	_ = json.NewDecoder(r.Body).Decode(&product)

	product.Uuid = uuid.New().String()

	result, err := collection.InsertOne(context.TODO(), product)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)

}
