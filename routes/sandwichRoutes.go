package routes

import (
	"github.com/gofiber/fiber/v2"
	"sandwhich/configs"
	"sandwhich/controllers"
	"sandwhich/middlewares"
)

func SandwichRoutes(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware(configs.GetEnv("SECRET"))

	router.Get("/", jwt, middlewares.EnsureUser, controllers.GetSandwiches)
	router.Get("/:id", jwt, middlewares.EnsureUser, controllers.GetSandwich)
	router.Get("/user/:id", jwt, middlewares.EnsureUser, controllers.GetSandwichesByUserID)
	router.Post("/create", jwt, middlewares.EnsureUser, controllers.CreateSandwich)
	router.Put("/update", jwt, middlewares.EnsureUser, controllers.UpdateSandwich)
}
