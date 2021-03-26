package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Book model
type Book struct {
	ID			primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn		string							`json:"isbn,omitempty" bson:"isbn,omitempty"`
	Title		string							`json:"title" bson:"title,omitempty"`
	Author	*Author							`json:"author" bson:"author,omitempty"`
}

// Author model
type Author struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

// Product Model
type Product struct {
	ID					primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Title				string							`json:"title,omitempty" bson:"title,omitempty"`
	Price				string							`json:"price" bson:"price,omitempty"`
	Category		string							`json:"category" bson:"category,omitempty"`
	Description	string							`json:"description" bson:"description,omitempty"`
	Image				string							`json:"image" bson:"image,omitempty"`
	Uuid				string							`json:"uuid" bson:"uuid,omitempty"`
}
