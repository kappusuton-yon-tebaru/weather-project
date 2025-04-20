//go:build wireinject
// +build wireinject

package app

import (
	"github.com/Porsche-ths/weather-4-u/geolocation-service/internal/api/geocode"
	"github.com/Porsche-ths/weather-4-u/geolocation-service/internal/api/geoip"
	"github.com/Porsche-ths/weather-4-u/geolocation-service/internal/config"
	"github.com/Porsche-ths/weather-4-u/geolocation-service/internal/geoapify"
	"github.com/google/wire"
)

func Intialize() (*App, error) {
	wire.Build(
		config.Load,
		geoapify.NewClient,
		geocode.NewService,
		geocode.NewHandler,
		geoip.NewService,
		geoip.NewHandler,
		New,
	)

	return &App{}, nil
}
