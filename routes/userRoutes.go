package routes

import (
	"github.com/gofiber/fiber/v2"
	"sandwhich/controllers"
)

func UserRoutes(router fiber.Router) {
	router.Get("/", controllers.GetUsers)
}
