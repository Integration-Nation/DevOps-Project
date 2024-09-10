package routes

import "github.com/gofiber/fiber"

func apiRoutes(app *fiber.App, ac controllers) {
	api:= app.Group("/api")
	api.Get("/search")
}