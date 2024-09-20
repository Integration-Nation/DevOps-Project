package controllers

import (
	"DevOps-Project/internal/services"

	"github.com/gofiber/fiber/v2"
)

type WeatherControllerI interface {
	GetWeather(c *fiber.Ctx) error
}

type WeatherController struct {
	service services.WeatherServiceI
}


func NewWeatherController(service services.WeatherServiceI) *WeatherController {
	return &WeatherController{service: service}
}

func (wc *WeatherController) GetWeather(c *fiber.Ctx) error {

	// Get the weather data
	weather, err := wc.service.GetWeather()
	if err != nil {
		// Return an internal server error status with an error message in JSON
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// On success, return a JSON response with the weather data
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"weather": weather,
	})
}