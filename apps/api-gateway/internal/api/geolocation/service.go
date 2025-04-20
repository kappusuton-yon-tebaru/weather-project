package geolocation

import (
	"encoding/json"
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

func (s *Service) GetLocation(reqQuery models.RequestParams) (models.Geolocation, error) {
	endpoint := ""
	if reqQuery.IpAddress != "" {
		endpoint = s.config.GeolocationServiceEndpoint + "/ip?ip=" + reqQuery.IpAddress

		resp, err := http.Get(endpoint)
		if err != nil {
			return models.Geolocation{}, err
		}
		defer resp.Body.Close()

		geoResponse := models.GeolocationIPResponse{}
		if err := json.NewDecoder(resp.Body).Decode(&geoResponse); err != nil {
			return models.Geolocation{}, err
		}

		return geoResponse.Result, nil
	} else {
		endpoint = s.config.GeolocationServiceEndpoint + "/search?text=" + reqQuery.Address

		resp, err := http.Get(endpoint)
		if err != nil {
			return models.Geolocation{}, err
		}
		defer resp.Body.Close()

		geoResponse := models.GeolocationAddressResponse{}
		if err := json.NewDecoder(resp.Body).Decode(&geoResponse); err != nil {
			return models.Geolocation{}, err
		}

		if len(geoResponse.Results) == 0 {
			return models.Geolocation{}, nil
		}

		return geoResponse.Results[0], nil
	}
}
