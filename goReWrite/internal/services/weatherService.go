package services

import (
	"DevOps-Project/internal/models"
	"go.uber.org/zap"
	"DevOps-Project/internal/utilities"
	
)

type WeatherServiceI interface {
	GetWeather() (*models.Weather, error)
}

type WeatherService struct {
	logger *zap.Logger

}

func NewWeatherService() *WeatherService {
	return &WeatherService{logger: utilities.NewLogger()}
}

func (ws *WeatherService) GetWeather() (*models.Weather, error) {
	// Lav API kaldet her
weather := &models.Weather{
		Temperature: 20,
		Condition:   "Sunny",
	}
	return weather, nil
}