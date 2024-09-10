package main

import (
	"DevOps-Project/internal/initializers"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnv() 
	initializers.ConnectDB()
}

func main() {
	app := fiber.New()
	app.Use(cors.New())




	
	app.Get("/api/search", GetSearch)


	err:= app.listen(":9090")
	if err!= nil {
		panic(err)
	}
}