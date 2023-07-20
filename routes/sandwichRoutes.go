package routes

import (
	"github.com/gofiber/fiber/v2"
	"sandwhich/controllers"
	"sandwhich/middlewares"
)

func SandwichRoutes(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware("secret")

	router.Get("/", jwt, controllers.GetSandwiches)
	router.Get("/:id", jwt, controllers.GetSandwich)
	router.Get("/user/:id", jwt, controllers.GetSandwichesByUserID)
	router.Post("/create", jwt, controllers.CreateSandwich)
	router.Put("/update", jwt, controllers.UpdateSandwich)
}
