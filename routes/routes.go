package routes

import (
	"my-fiber-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/home", controllers.GetHome)
	app.Get("/about", controllers.GetAbout)
	app.Get("/process", controllers.GetProcess)
	app.Get("/portfolio", controllers.GetPortfolio)
}
