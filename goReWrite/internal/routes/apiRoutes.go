package routes

import (
	"DevOps-Project/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func PageRoutes(app *fiber.App, pc controllers.PageControllerI) {
	api := app.Group("/api")
	api.Post("/search", pc.GetSearchResults)

}

func UserRoutes(app *fiber.App, uc controllers.UserControllerI) {
	api := app.Group("/api")
	api.Post("/register", uc.Register)
	api.Post("/login", uc.Login)
	//api.Post("/logout", uc.Logout)
	api.Get("/users", uc.GetAllUsers)
}
