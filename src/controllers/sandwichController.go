package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sandwhich/src/configs"
	models2 "sandwhich/src/models"
	"time"
)

var sandwichCollection = configs.GetCollection(configs.DB, "sandwiches")

func GetSandwiches(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	limit := 100
	cursor, err := sandwichCollection.Find(ctx, bson.M{}, options.Find().SetLimit(int64(limit)))
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	var sandwiches []models2.Sandwich

	for cursor.Next(ctx) {
		var sandwich models2.Sandwich
		if err := cursor.Decode(&sandwich); err != nil {
			return err
		}
		sandwiches = append(sandwiches, sandwich)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return c.JSON(sandwiches)
}

func GetSandwich(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)

	var sandwich models2.Sandwich
	filter := bson.M{"_id": id}
	err = sandwichCollection.FindOne(ctx, filter).Decode(&sandwich)
	if err != nil {
		return err
	}

	return c.JSON(sandwich)
}

func GetSandwichesByUserID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	cursor, err := sandwichCollection.Find(ctx, bson.M{"userId": id})

	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	var sandwiches []models2.Sandwich

	for cursor.Next(ctx) {
		var sandwich models2.Sandwich
		if err := cursor.Decode(&sandwich); err != nil {
			return err
		}
		sandwiches = append(sandwiches, sandwich)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return c.JSON(sandwiches)
}

func CreateSandwich(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := c.Locals("user").(*models2.User)

	sandwich := new(models2.Sandwich)
	if err := c.BodyParser(sandwich); err != nil {
		return err
	}

	if sandwich.Name == "" || sandwich.UserId.IsZero() || sandwich.ImageUrl == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	if sandwich.UserId != user.Id {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "UserId does not refer to the connected user",
		})
	}

	sandwich.Id = primitive.NewObjectID()
	result, err := sandwichCollection.InsertOne(ctx, sandwich)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func UpdateSandwich(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := c.Locals("user").(*models2.User)

	sandwichUpdated := new(models2.Sandwich)
	if err := c.BodyParser(sandwichUpdated); err != nil {
		return err
	}

	// get sandwich from mongo and check if user is owner
	var sandwich models2.Sandwich
	filter := bson.M{"_id": sandwichUpdated.Id}
	err := sandwichCollection.FindOne(ctx, filter).Decode(&sandwich)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Sandwich not found",
		})
	}
	if sandwich.UserId != user.Id {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	if sandwichUpdated.Name == "" || sandwichUpdated.UserId.IsZero() || sandwichUpdated.ImageUrl == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	res, err := sandwichCollection.UpdateOne(ctx, bson.M{"_id": sandwichUpdated.Id}, bson.M{"$set": sandwichUpdated})
	if err != nil {
		return err
	}

	return c.JSON(res)
}
