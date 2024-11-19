// @title whoKnows-goFiber
// @version 0.5.0
// @description This is a search engine API for managing users and pages.
// @host localhost:9090
// @BasePath /

package main

import (
	"DevOps-Project/internal/controllers"
	"DevOps-Project/internal/initializers"
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/monitoring"
	"DevOps-Project/internal/repositories"
	"DevOps-Project/internal/routes"
	"DevOps-Project/internal/services"
	"os"

	"log"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	//initializers.ConnectSqlite()
}

func main() {

	if err := initializers.DB.AutoMigrate(&models.Page{}); err != nil {
		log.Fatal(err)
	}

	if err := initializers.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}

	// initializers.MigrateUsers()
	//initializers.MigratePages()

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	app := fiber.New()
	app.Use(cors.New())

	go monitoring.CollectSystemMetrics()

	prometheus := fiberprometheus.New("WhoKnows-goFiber")
	prometheus.RegisterAt(app, "/metrics")
	prometheus.SetSkipPaths([]string{"/ping"})
	app.Use(prometheus.Middleware)

	app.Use(func(c *fiber.Ctx) error {
		monitoring.ConcurrentRequests.Inc()       // Increment concurrent requests
		defer monitoring.ConcurrentRequests.Dec() // Decrement when done

		return c.Next()
	})

	v := validator.New()

	pageRepo := repositories.NewPageRepository(initializers.DB)
	pageService := services.NewPageService(pageRepo)
	pageController := controllers.NewPageController(pageService, v)

	userRepo := repositories.NewUserRepository(initializers.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService, v)

	weatherService := services.NewWeatherService()
	weatherController := controllers.NewWeatherController(weatherService)

	routes.PageRoutes(app, pageController)
	routes.WeatherRoutes(app, weatherController)
	routes.UserRoutes(app, userController, jwtSecret)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})


	//Start HTTPS server
// 	err := app.ListenTLS(":9090", "/etc/letsencrypt/live/integration-nation.dk/fullchain.pem", "/etc/letsencrypt/live/integration-nation.dk/privkey.pem")
  
// 	if err != nil {
// 		panic(err)
// 	}
// }

	err := app.Listen(":9090")
	if err != nil {
		panic(err)
	}
}
