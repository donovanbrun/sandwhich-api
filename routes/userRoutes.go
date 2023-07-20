package routes

import (
	"github.com/gofiber/fiber/v2"
	"sandwhich/configs"
	"sandwhich/controllers"
	"sandwhich/middlewares"
)

func UserRoutes(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware(configs.GetEnv("SECRET"))

	router.Get("/find-admin/:id", jwt, middlewares.EnsureAdmin, controllers.GetUser)
	router.Get("/find/:id", jwt, middlewares.EnsureUser, controllers.GetUserPublic)
	router.Post("/login", controllers.Login)
	router.Post("/signup", controllers.Signup)
	router.Post("/update", jwt, middlewares.EnsureUser, controllers.UpdateUser)
}
