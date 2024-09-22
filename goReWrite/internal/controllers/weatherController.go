package controllers

import (
	"DevOps-Project/internal/services"
	"DevOps-Project/internal/utilities"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type WeatherControllerI interface {
	GetWeather(c *fiber.Ctx) error
}

type WeatherController struct {
	service services.WeatherServiceI
	logger  *zap.Logger
}


func NewWeatherController(service services.WeatherServiceI) *WeatherController {
	return &WeatherController{
		service: service,
		logger: utilities.NewLogger(),
	}
}

func (wc *WeatherController) GetWeather(c *fiber.Ctx) error {

	// Get the weather data
	weather, err := wc.service.GetWeather()
	if err != nil {
		// Return an internal server error status with an error message in JSON
		wc.logger.Error("Failed to get weather data", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// On success, return a JSON response with the weather data
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"weather": weather,
	})
}