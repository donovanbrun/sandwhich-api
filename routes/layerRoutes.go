package routes

import (
	"github.com/gofiber/fiber/v2"
	"sandwhich/configs"
	"sandwhich/controllers"
	"sandwhich/middlewares"
)

func LayerRoute(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware(configs.GetEnv("SECRET"))
	router.Get("/", jwt, middlewares.EnsureUser, controllers.GetLayers)
}
