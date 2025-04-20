//go:build wireinject
// +build wireinject

package app

import (
	"github.com/Porsche-ths/weather-4-u/weather-service/internal/api"
	"github.com/Porsche-ths/weather-4-u/weather-service/internal/config"
	"github.com/Porsche-ths/weather-4-u/weather-service/internal/client"
	"github.com/google/wire"
)

func Initialize() (*App, error) {
	wire.Build(
		config.Load,
		client.NewClient,
		api.NewService,
		api.NewHandler,
		New,
	)

	return &App{}, nil
}
