package main

import (
	"my-fiber-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Routing
	routes.SetupRoutes(app)

	// (opsional) serve file static
	app.Static("/", "./static")

	app.Listen(":3001")
}
