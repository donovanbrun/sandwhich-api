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

var userCollection = configs.GetCollection(configs.DB, "users")

func GetConnectedUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	userPublic := models.UserPublic{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
		Bio:   user.Bio,
	}
	return c.JSON(userPublic)
}

func GetUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)

	var user models.User
	filter := bson.M{"_id": id}
	err = userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func GetUserPublic(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)

	var user models.UserPublic
	filter := bson.M{"_id": id}
	err = userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func GetByCredentials(email, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	filter := bson.M{"email": email}
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	if !CheckPasswordHash(password, user.Password) {
		return nil, err
	}

	return &user, nil
}

func ExistByEmail(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	filter := bson.M{"email": email}
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return false
	}

	if (models.User{}) == user {
		return false
	}
	return true
}

func UpdateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userTDO models.User
	err := c.BodyParser(&userTDO)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": userTDO.Id}
	update := bson.M{"$set": bson.M{
		"name": userTDO.Name,
		"bio":  userTDO.Bio,
	}}

	res, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func GetUserById(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oid, _ := primitive.ObjectIDFromHex(id)
	var user models.User
	filter := bson.M{"_id": oid}
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
