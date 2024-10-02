package main

import (
	"DevOps-Project/internal/controllers"
	"DevOps-Project/internal/initializers"
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/repositories"
	"DevOps-Project/internal/routes"
	"DevOps-Project/internal/services"
	"os"
	"time"

	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartTokenCleanupScheduler(interval time.Duration, userRepo *repositories.TokenBlacklistRepository) {
	go func() {
		for {
			time.Sleep(interval)
			err := userRepo.CleanupExpiredTokens()
			if err != nil {
				log.Println("Error cleaning up expired tokens:", err)
			}
		}
	}()
}

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {

	if err := initializers.DB.AutoMigrate(&models.Page{}); err != nil {
		log.Fatal(err)
	}

	if err := initializers.DB.AutoMigrate(&models.TokenBlacklist{}); err != nil {
		log.Fatal(err)
	}

	if err := initializers.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	app := fiber.New()
	app.Use(cors.New())

	v := validator.New()

	pageRepo := repositories.NewPageRepository(initializers.DB)
	pageService := services.NewPageService(pageRepo)
	pageController := controllers.NewPageController(pageService, v)

	blackListRepo := repositories.NewTokenBlacklistRepository(initializers.DB)

	userRepo := repositories.NewUserRepository(initializers.DB)
	userService := services.NewUserService(userRepo, blackListRepo)
	userController := controllers.NewUserController(userService, v)

	weatherService := services.NewWeatherService()
	weatherController := controllers.NewWeatherController(weatherService)

	routes.PageRoutes(app, pageController)
	routes.WeatherRoutes(app, weatherController)
	routes.UserRoutes(app, userController, jwtSecret)

	go StartTokenCleanupScheduler(1*time.Hour, blackListRepo)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start HTTPS server
	// err := app.ListenTLS(":9090", "/etc/letsencrypt/live/40-85-136-203.nip.io/fullchain.pem", "/etc/letsencrypt/live/40-85-136-203.nip.io/privkey.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := app.Listen(":9090")
	if err != nil {
		panic(err)
	}
}
