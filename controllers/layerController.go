package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"sandwhich/configs"
	"sandwhich/models"
	"time"
)

var layerCollection = configs.GetCollection(configs.DB, "layers")

func GetLayers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := layerCollection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	var layers []models.Layer

	for cursor.Next(ctx) {
		var layer models.Layer
		if err := cursor.Decode(&layer); err != nil {
			return err
		}
		layers = append(layers, layer)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return c.JSON(layers)
}
