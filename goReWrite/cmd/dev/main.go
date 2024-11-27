package main

// @title whoKnows-goFiber
// @version 0.5.0
// @description This is a search engine API for managing users and pages.
// @host localhost:9090
// @BasePath /

import (
	"DevOps-Project/internal/controllers"
	"DevOps-Project/internal/initializers"
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/monitoring"
	"DevOps-Project/internal/repositories"
	"DevOps-Project/internal/routes"
	"DevOps-Project/internal/services"
	"DevOps-Project/internal/utilities"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.ConnectSqliteLocal()
}

func prometheusHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		handler := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
		handler(c.Context())
		return nil
	}
}

func main() {
	if err := initializers.DB.AutoMigrate(&models.Page{}); err != nil {
		log.Fatal(err)
	}

	if err := initializers.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}

	//initializers.MigrateUsers()
	//initializers.MigratePages()

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	app := fiber.New()
	app.Use(cors.New())

	go monitoring.CollectSystemMetrics()

	// Middleware for tracking concurrent requests and errors
	app.Use(func(c *fiber.Ctx) error {
		monitoring.HTTPConcurrentRequests.Inc()
		defer monitoring.HTTPConcurrentRequests.Dec()
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			monitoring.HTTPTotalErrors.Inc()
			return err
		}
		return nil
	})

	app.Get("/metrics", prometheusHandler())
	v := validator.New()
	logger := utilities.NewMockLogger()

	pageRepo := repositories.NewPageRepository(initializers.DB, logger)
	pageService := services.NewPageService(pageRepo, logger)
	pageController := controllers.NewPageController(pageService, v, logger)

	userRepo := repositories.NewUserRepository(initializers.DB, logger)
	userService := services.NewUserService(userRepo, logger)
	userController := controllers.NewUserController(userService, v, logger)

	weatherService := services.NewWeatherService(logger)
	weatherController := controllers.NewWeatherController(weatherService, logger)

	routes.PageRoutes(app, pageController)
	routes.WeatherRoutes(app, weatherController)
	routes.UserRoutes(app, userController, jwtSecret)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":7070")
	if err != nil {
		panic(err)
	}
}
