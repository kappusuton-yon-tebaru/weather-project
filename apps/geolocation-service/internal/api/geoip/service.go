package geoip

import (
	"context"
	"fmt"

	"github.com/Porsche-ths/weather-4-u/geolocation-service/internal/geoapify"
	"github.com/Porsche-ths/weather-4-u/geolocation-service/internal/models"
)

type Service struct {
	geoapifyClient *geoapify.Client
}

func NewService(geoapifyClient *geoapify.Client) *Service {
	return &Service{
		geoapifyClient: geoapifyClient,
	}
}

func (svc *Service) CoordinateFromIP(ctx context.Context, ip string) (models.Location, error) {
	resp, err := svc.geoapifyClient.CoordinateFromIP(ctx, ip)
	if err != nil {
		return models.Location{}, err
	}

	location := models.Location{
		Coordinate: resp.Location,
		Country:    resp.Country.Name,
		City:       resp.City.Name,
		Address:    fmt.Sprintf("%s, %s", resp.City.Name, resp.Country.Name),
	}

	return location, nil
}
