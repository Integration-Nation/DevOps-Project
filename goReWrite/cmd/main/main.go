package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {

}

func main() {
	app := fiber.New()
	app.Use(cors.New())




	
	app.Get("/api/search",getSearch())


	err:= app.listen(":9090")
	if err!= nil {
		panic(err)
	}
}