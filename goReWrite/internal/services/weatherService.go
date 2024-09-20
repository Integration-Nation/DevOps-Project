package services

import (
	"DevOps-Project/internal/models"
)

type WeatherServiceI interface {
	GetWeather() (*models.Weather, error)
}

type WeatherService struct {
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (ws *WeatherService) GetWeather() (*models.Weather, error) {
	// Lav API kaldet her
weather := &models.Weather{
		Temperature: 20,
		Condition:   "Sunny",
	}
	return weather, nil
}