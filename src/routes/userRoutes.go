package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers2 "sandwhich/src/controllers"
	"sandwhich/src/middlewares"
)

func UserRoutes(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware()
	filter := middlewares.AuthorizedSource()

	router.Get("", filter, jwt, middlewares.EnsureUser, controllers2.GetConnectedUser)
	router.Get("/find-admin/:id", filter, jwt, middlewares.EnsureAdmin, controllers2.GetUser)
	router.Get("/find/:id", controllers2.GetUserPublic)
	router.Post("/login", filter, controllers2.Login)
	router.Post("/signup", filter, controllers2.Signup)
	router.Put("/update", filter, jwt, middlewares.EnsureUser, controllers2.UpdateUser)
}
