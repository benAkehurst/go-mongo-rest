package dogs_handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"quickstart/helper"
)

// This is the structure that comes back from the api
type ApiResponse struct {
	Message		string		`json:"message"`
}

// This is the format that needs to be returned to the user
type FormattedResponse struct {
	Url		string		`json:"url"`
}

// var collection = helper.ConnectDB().Collection("dogs")

func GetRandomDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Calls the external api and stores the response
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")
		if err != nil {
		helper.GetError(err, w)
		return
	}

	// The response body is read
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		helper.GetError(err, w)
		return
	}

	var responseObject ApiResponse
	// convert the response data to match the expected structure
  json.Unmarshal(responseData, &responseObject)
	// data is then encoded in proper json response
	json.NewEncoder(w).Encode(
		&FormattedResponse{Url: responseObject.Message},
	)
}
