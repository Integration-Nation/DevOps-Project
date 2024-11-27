package controllers

import (
	"DevOps-Project/internal/services"

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

func NewWeatherController(service services.WeatherServiceI, logger *zap.Logger) *WeatherController {
	return &WeatherController{
		service: service,
		logger:  logger,
	}
}

// GetWeather godoc
// @Summary Get weather data for a given location
// @Description Fetch weather information based on latitude and longitude. Defaults to Copenhagen if no query parameters are provided.
// @Tags weather
// @Accept json
// @Produce json
// @Param latitude query string false "Latitude of the location" default(55.6761)
// @Param longitude query string false "Longitude of the location" default(12.5683)
// @Success 200 {object} map[string]interface{} "Returns weather data"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /weather [get]
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
