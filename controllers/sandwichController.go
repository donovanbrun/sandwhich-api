package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sandwhich/configs"
	"sandwhich/models"
	"time"
)

var sandwichCollection = configs.GetCollection(configs.DB, "sandwiches")

func GetSandwiches(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := sandwichCollection.Find(context, bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context)

	var sandwiches []models.Sandwich

	for cursor.Next(context) {
		var sandwich models.Sandwich
		if err := cursor.Decode(&sandwich); err != nil {
			return err
		}
		sandwiches = append(sandwiches, sandwich)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return ctx.JSON(sandwiches)
}

func GetSandwich(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idParam := ctx.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	filter := bson.M{"_id": id}

	var sandwich models.Sandwich
	err = sandwichCollection.FindOne(context, filter).Decode(&sandwich)
	if err != nil {
		return err
	}

	return ctx.JSON(sandwich)
}
