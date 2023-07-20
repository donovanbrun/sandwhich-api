package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Bio      string             `json:"bio" bson:"bio"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Admin    bool               `json:"admin" bson:"admin"`
}

type UserPublic struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Bio  string             `json:"bio" bson:"bio"`
}
