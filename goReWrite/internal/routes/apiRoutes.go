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
	api.Get("/logout", uc.Logout)
	api.Get("/users", uc.GetAllUsers)
}

func WeatherRoutes(app *fiber.App, wc controllers.WeatherControllerI) {
		api := app.Group("/api")
	api.Get("/weather", wc.GetWeather)
}
