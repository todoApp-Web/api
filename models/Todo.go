package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
    ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Text    string           `json:"text,omitempty" bson:"text,omitempty"`
	Completed bool           `json:"completed,omitempty" bson:"completed,omitempty"`
}