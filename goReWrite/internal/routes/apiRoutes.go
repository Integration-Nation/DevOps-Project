package routes

import (
	"DevOps-Project/internal/controllers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func PageRoutes(app *fiber.App, pc controllers.PageControllerI) {
	api := app.Group("/api")
	api.Get("/search", pc.GetSearchResults)

}

func UserRoutes(app *fiber.App, uc controllers.UserControllerI, jwtSecret string) {
	api := app.Group("/api")
	api.Post("/register", uc.Register)
	api.Post("/login", uc.Login)

	app.Use("/api", jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
	}))

	api.Get("/logout", uc.Logout)
	api.Get("/users", uc.GetAllUsers)
	api.Delete("/user", uc.DeleteUser)

}

func WeatherRoutes(app *fiber.App, wc controllers.WeatherControllerI) {
	api := app.Group("/api")
	api.Get("/weather", wc.GetWeather)
}
