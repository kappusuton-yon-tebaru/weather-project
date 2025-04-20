package geoapify

import "github.com/Porsche-ths/weather-4-u/geolocation-service/internal/models"

type ErrorResponse struct {
	Message string `json:"message"`
}

type GeocodeResponse struct {
	Features []feature `json:"features"`
}

type feature struct {
	Property property `json:"properties"`
}

type property struct {
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Lon       float32 `json:"lon"`
	Lat       float32 `json:"lat"`
	Formatted string  `json:"formatted"`
}

type GeoipResponse struct {
	City     city              `json:"city"`
	Country  country           `json:"country"`
	Location models.Coordinate `json:"location"`
}

type city struct {
	Name string `json:"name"`
}

type country struct {
	Name string `json:"name"`
}
