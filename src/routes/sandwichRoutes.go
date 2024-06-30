package routes

import (
	"github.com/gofiber/fiber/v2"
	"sandwhich/src/controllers"
	"sandwhich/src/middlewares"
)

func SandwichRoutes(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware()

	router.Get("/", controllers.GetSandwiches)
	router.Get("/:id", controllers.GetSandwich)
	router.Get("/user/:id", jwt, middlewares.EnsureUser, controllers.GetSandwichesByUserID)
	router.Post("/create", jwt, middlewares.EnsureUser, controllers.CreateSandwich)
	router.Put("/update", jwt, middlewares.EnsureUser, controllers.UpdateSandwich)
}
