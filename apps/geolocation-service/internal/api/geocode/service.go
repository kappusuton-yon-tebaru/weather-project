package geocode

import (
	"context"

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

func (svc *Service) SearchFromAddress(ctx context.Context, text string) ([]models.Location, error) {
	resp, err := svc.geoapifyClient.SearchFromAddress(ctx, text)
	if err != nil {
		return []models.Location{}, err
	}

	locations := make([]models.Location, 0, len(resp.Features))

	for _, feat := range resp.Features {
		locations = append(locations, models.Location{
			Country: feat.Property.Country,
			City:    feat.Property.City,
			Address: feat.Property.Formatted,
			Coordinate: models.Coordinate{
				Latitude:  feat.Property.Lat,
				Longitude: feat.Property.Lon,
			},
		})
	}

	return locations, nil
}
