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
		logger:  utilities.NewLogger(),
	}
}

func (wc *WeatherController) GetWeather(c *fiber.Ctx) error {
	defaultLatitude := "55.6761"
	defaultLongitude := "12.5683"

	latitude := c.Query("latitude", defaultLatitude)
	longitude := c.Query("longitude", defaultLongitude)

	weather, err := wc.service.GetWeather(latitude, longitude)
	if err != nil {

		wc.logger.Error("Failed to get weather data", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"weather": weather,
	})
}
