package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers2 "sandwhich/src/controllers"
	"sandwhich/src/middlewares"
)

func UserRoutes(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware()

	router.Get("", jwt, middlewares.EnsureUser, controllers2.GetConnectedUser)
	router.Get("/find-admin/:id", jwt, middlewares.EnsureAdmin, controllers2.GetUser)
	router.Get("/find/:id", jwt, middlewares.EnsureUser, controllers2.GetUserPublic)
	router.Post("/login", controllers2.Login)
	router.Post("/signup", controllers2.Signup)
	router.Post("/update", jwt, middlewares.EnsureUser, controllers2.UpdateUser)
}
