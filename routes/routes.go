package routes

import "github.com/gofiber/fiber/v2"

func InitRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	initUserRoutes(api)
}
