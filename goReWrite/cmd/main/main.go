package main

import (
	"DevOps-Project/internal/controllers"
	"DevOps-Project/internal/initializers"
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/repositories"
	"DevOps-Project/internal/routes"
	"DevOps-Project/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {

	initializers.DB.AutoMigrate(&models.Page{})
	initializers.DB.AutoMigrate(&models.User{})

	    jwtSecret := os.Getenv("JWT_SECRET")
        if jwtSecret == "" {
        log.Fatal("JWT_SECRET environment variable not set")
    }

	app := fiber.New()
	app.Use(cors.New())

	v := validator.New()

     app.Use("/api", jwtware.New(jwtware.Config{
        SigningKey: []byte(jwtSecret),
    }))

	pageRepo := repositories.NewPageRepository(initializers.DB)
	pageService := services.NewPageService(pageRepo)
	pageController := controllers.NewPageController(pageService, v)

	userRepo := repositories.NewUserRepository(initializers.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService, v)

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
