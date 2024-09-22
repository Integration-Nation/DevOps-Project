package main

import (
	"DevOps-Project/internal/controllers"
	"DevOps-Project/internal/initializers"
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/repositories"
	"DevOps-Project/internal/routes"
	"DevOps-Project/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {

	initializers.DB.AutoMigrate(&models.Page{})
	initializers.DB.AutoMigrate(&models.User{})

	app := fiber.New()
	app.Use(cors.New())

	pageRepo := repositories.NewPageRepository(initializers.DB)
	pageService := services.NewPageService(pageRepo)
	pageController := controllers.NewPageController(pageService)

	userRepo := repositories.NewUserRepository(initializers.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	weatherService := services.NewWeatherService()
	weatherController := controllers.NewWeatherController(weatherService)

	routes.PageRoutes(app, pageController)
	routes.UserRoutes(app, userController)
	routes.WeatherRoutes(app, weatherController)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":9090")
	if err != nil {
		panic(err)
	}
}