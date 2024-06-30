package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"sandwhich/src/configs"
	"sandwhich/src/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization,Access-Control-Allow-Headers,Access-Control-Allow-Methods,Access-Control-Allow-Credentials",
		AllowOrigins:     "https://sandwhich.donovanbrun.com,http://localhost:3000",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	configs.ConnectDB()

	routes.Routes(app)

	port := configs.Env.PORT
	if port == "" {
		port = "8080"
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("UP")
	})

	app.Listen("0.0.0.0:" + port)
}
