package dto

import "github.com/Porsche-ths/weather-4-u/weather-service/internal/models"

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

type WeatherResponse struct {
    Location	models.Location		`json:"location"`
	Current		models.Current		`json:"current"`
	Forecast	models.Forecast		`json:"forecast"`	
}
