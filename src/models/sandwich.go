package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sandwich struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
	Description string             `json:"description" bson:"description"`
	Layers      []string           `json:"layers" bson:"layers"`
	ImageUrl    string             `json:"imageUrl" bson:"imageUrl"`
}
