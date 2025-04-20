package app

import (
	"github.com/Porsche-ths/weather-4-u/geolocation-service/internal/api/geocode"
	"github.com/Porsche-ths/weather-4-u/geolocation-service/internal/api/geoip"
	"github.com/Porsche-ths/weather-4-u/geolocation-service/internal/config"
)

type App struct {
	Config         *config.Config
	GeocodeHandler *geocode.Handler
	GeoipHandler   *geoip.Handler
}

func New(
	Config *config.Config,
	GeocodeHandler *geocode.Handler,
	GeoipHandler *geoip.Handler,
) *App {
	return &App{
		Config,
		GeocodeHandler,
		GeoipHandler,
	}
}
