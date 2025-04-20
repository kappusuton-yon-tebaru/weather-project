package weather

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Porsche-ths/weather-4-u/api-gateway/internal/config"
	"github.com/Porsche-ths/weather-4-u/api-gateway/internal/models"
)

type Service struct {
	config *config.Config
}

func NewService(config *config.Config) *Service {
	return &Service{
		config: config,
	}
}

func (s *Service) GetWeather(coordinate models.Coordinate) (models.WeatherResponse, error) {
	endpoint := ""
	endpoint = s.config.WeatherServiceEndpoint + "/weather?q=" + fmt.Sprintf("%f", coordinate.Latitude) + "," + fmt.Sprintf("%f", coordinate.Longitude) + "&days=1"

	resp, err := http.Get(endpoint)
	if err != nil {
		return models.WeatherResponse{}, err
	}
	defer resp.Body.Close()

	weatherResponse := models.WeatherResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return models.WeatherResponse{}, err
	}

	return weatherResponse, nil
}
