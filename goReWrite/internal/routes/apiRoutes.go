package routes

import (
	"DevOps-Project/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App, pc controllers.PageControllerI) {
	api := app.Group("/api")
	api.Post("/search", pc.GetSearchResults)

}
