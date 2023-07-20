package routes

import (
	"github.com/gofiber/fiber/v2"
	"sandwhich/controllers"
	"sandwhich/middlewares"
)

func UserRoutes(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware("secret")

	//router.Get("/find/:id", jwt, controllers.GetUser)
	//router.Get("/find-public/:id", jwt, controllers.GetUserPublic)
	router.Get("/find/:id", jwt, controllers.GetUserPublic)
	router.Post("/login", controllers.Login)
	router.Post("/signup", controllers.Signup)
}
