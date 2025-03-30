package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Book - структура книги
type Book struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title  string             `bson:"title" json:"title"`
	Author string             `bson:"author" json:"author"`
}
