package routes

import (
	"github.com/gofiber/fiber/v2"
	"sandwhich/controllers"
)

func SandwichRoutes(router fiber.Router) {
	router.Get("/", controllers.GetSandwiches)
	router.Get("/:id", controllers.GetSandwich)
}
