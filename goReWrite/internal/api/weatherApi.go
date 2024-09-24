package api

import (
	"DevOps-Project/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func WeatherApi(latitude, longitude string) (*models.Weather, error) {

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current=temperature_2m,wind_speed_10m&hourly=temperature_2m", latitude, longitude)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	var weatherResponse models.Weather
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return nil, err
	}

	return &weatherResponse, nil
}
