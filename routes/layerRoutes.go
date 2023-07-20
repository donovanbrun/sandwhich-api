package routes

import (
	"github.com/gofiber/fiber/v2"
	"sandwhich/controllers"
)

func LayerRoute(router fiber.Router) {
	router.Get("/", controllers.GetLayers)
}
