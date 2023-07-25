package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	jtoken "github.com/golang-jwt/jwt/v4"
	"sandwhich/src/configs"
	"sandwhich/src/controllers"
)

// Middleware JWT function
func NewAuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(configs.Env.SECRET),
		ContextKey: "token",
	})
}

func EnsureUser(c *fiber.Ctx) error {
	token := c.Locals("token").(*jtoken.Token)
	if token == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	userId := token.Claims.(jtoken.MapClaims)["id"]

	user, err := controllers.GetUserById(userId.(string))
	if err != nil {
		return err
	}

	c.Locals("user", user)

	return c.Next()
}

func EnsureAdmin(c *fiber.Ctx) error {
	token := c.Locals("token").(*jtoken.Token)
	if token == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	userId := token.Claims.(jtoken.MapClaims)["id"]

	user, err := controllers.GetUserById(userId.(string))
	if err != nil {
		return err
	}

	c.Locals("user", user)

	if !user.Admin {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return c.Next()
}
