package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"sandwhich/configs"
	"sandwhich/models"
	"time"
)

var userCollection = configs.GetCollection(configs.DB, "users")

func GetUsers(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := userCollection.Find(context, bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context)

	var users []models.User

	for cursor.Next(context) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return ctx.JSON(users)
}
