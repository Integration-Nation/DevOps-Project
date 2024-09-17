package routes

import (
	"DevOps-Project/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	api:= app.Group("/api")
	api.Get("/search", controllers.GetSearchResults)

}