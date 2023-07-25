package routes

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	api := app.Group("/api")
	api.Route("/sandwich", SandwichRoutes)
	api.Route("/user", UserRoutes)
	api.Route("/layer", LayerRoute)
}
