package main

import (
	"github.com/gofiber/fiber/v2"
	"sandwhich/configs"
	"sandwhich/routes"
)

func main() {
	app := fiber.New()
	configs.ConnectDB()

	routes.Routes(app)

	app.Listen("0.0.0.0:1234")
}
