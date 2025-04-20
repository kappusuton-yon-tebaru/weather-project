package app

import (
	"github.com/Porsche-ths/weather-4-u/weather-service/internal/api"
	"github.com/Porsche-ths/weather-4-u/weather-service/internal/config"
)

type App struct {
	Config         	*config.Config
	WeatherHandler	*api.Handler
}

func New(
	Config *config.Config,
	WeatherHandler *api.Handler,
) *App {
	return &App{
		Config,
		WeatherHandler,
	}
}
