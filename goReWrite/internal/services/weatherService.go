package services

import (
	"DevOps-Project/internal/api"
	"DevOps-Project/internal/models"

	"go.uber.org/zap"
)

type WeatherServiceI interface {
	GetWeather(latitude, longitude string) (*models.Weather, error)
}

type WeatherService struct {
	logger *zap.Logger
}

func NewWeatherService(logger *zap.Logger) *WeatherService {
	return &WeatherService{logger: logger}
}

func (ws *WeatherService) GetWeather(latitude, longitude string) (*models.Weather, error) {
	// Lav API kaldet her
	weather, err := api.WeatherApi(latitude, longitude)

	if err != nil {
		return nil, err
	}

	return weather, nil
}
