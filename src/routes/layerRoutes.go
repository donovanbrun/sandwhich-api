package routes

import (
	"github.com/gofiber/fiber/v2"
	"sandwhich/src/controllers"
	"sandwhich/src/middlewares"
)

func LayerRoute(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware()
	router.Get("/", jwt, middlewares.EnsureUser, controllers.GetLayers)
}
