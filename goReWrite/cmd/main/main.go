package main

import (
	"DevOps-Project/internal/initializers"
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {

	initializers.DB.AutoMigrate(&models.Page{})

	app := fiber.New()
	app.Use(cors.New())

	routes.ApiRoutes(app)

	err := app.Listen(":9090")
	if err != nil {
		panic(err)
	}
}
