package app

import (
	"github.com/Porsche-ths/weather-4-u/api-gateway/internal/api/geolocation"
	"github.com/Porsche-ths/weather-4-u/api-gateway/internal/api/weather"
	"github.com/Porsche-ths/weather-4-u/api-gateway/internal/config"
)

type App struct {
	Config             *config.Config
	GeolocationHandler *geolocation.Handler
	WeatherHandler     *weather.Handler
}

func New(Config *config.Config, GeolocationHandler *geolocation.Handler, WeatherHandler *weather.Handler) *App {
	return &App{
		Config,
		GeolocationHandler,
		WeatherHandler,
	}
}
