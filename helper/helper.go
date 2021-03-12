package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Helper function to connect to MongoDB
func ConnectDB() *mongo.Database {
	config := GetConfiguration()
	// set client options
	clientOptions := options.Client().ApplyURI(config.ConnectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("go_rest_api")

	return collection
}

// ErrorResponse: model
type ErrorResponse struct {
			StatusCode		int			`json:"status"`
			ErrorMessage	string	`json:"message"`
}

// GetError: helper to prepare error model
func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode: http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

// Configuration model
type Configuration struct {
		Port								string
		ConnectionString		string
}

// GetConfiguration: method to popuplate the config
func GetConfiguration() Configuration {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error Loading .env file")
	}

	configuration := Configuration{
		os.Getenv("PORT"),
		os.Getenv("CONNECTION_STRING"),
	}

	serverRunning := "Server running on port " + configuration.Port
  fmt.Println(serverRunning)

	return configuration
}
