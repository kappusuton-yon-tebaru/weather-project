package api

import (
	"context"

	// "github.com/Porsche-ths/weather-4-u/weather-service/internal/geoapify"
	"github.com/Porsche-ths/weather-4-u/weather-service/internal/models"
	"github.com/Porsche-ths/weather-4-u/weather-service/internal/client"
)

type Service struct {
	weatherClient *client.Client
}

func NewService(weatherClient *client.Client) *Service {
	return &Service{
		weatherClient: weatherClient,
	}
}

// func (svc *Service) CoordinateFromIP(ctx context.Context, ip string) (models.Location, error) {
// 	resp, err := svc.geoapifyClient.CoordinateFromIP(ctx, ip)
// 	if err != nil {
		// return models.Location{}, err
// 	}

// 	location := models.Location{
// 		Coordinate: resp.Location,
// 		Country:    resp.Country.Name,
// 		City:       resp.City.Name,
// 		Address:    fmt.Sprintf("%s, %s", resp.City.Name, resp.Country.Name),
// 	}

// 	return location, nil
// }

func (svc *Service) FetchWeather(ctx context.Context, q string, days string) (models.WeatherResponse, error) {
	resp, err := svc.weatherClient.FetchWeather(ctx, q, days)
	if err != nil {
		return models.WeatherResponse{}, err
	}

	data := models.WeatherResponse{
		Location: resp.Location,
		Current: resp.Current,
		Forecast: resp.Forecast,
	}

	// location := models.Location{
	// 	Coordinate: resp.Location,
	// 	Country:    resp.Country.Name,
	// 	City:       resp.City.Name,
	// 	Address:    fmt.Sprintf("%s, %s", resp.City.Name, resp.Country.Name),
	// }

	return data, nil
}
