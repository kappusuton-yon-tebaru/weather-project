package main

import (
	"fmt"

	"github.com/Porsche-ths/weather-4-u/api-gateway/cmd/app"
	"github.com/Porsche-ths/weather-4-u/api-gateway/internal/api/geolocation"
	"github.com/Porsche-ths/weather-4-u/api-gateway/internal/api/weather"
	"github.com/Porsche-ths/weather-4-u/api-gateway/internal/config"
	"github.com/Porsche-ths/weather-4-u/api-gateway/internal/router"
)

func main() {
	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	geolocationService := geolocation.NewService(config)
	weatherService := weather.NewService(config)

	geolocationHandler := geolocation.NewHandler(geolocationService, weatherService)
	weatherHandler := weather.NewHandler(weatherService)

	app := app.New(config, geolocationHandler, weatherHandler)

	r := router.New()
	r.RegisterRoutes(app)

	hostname := ""
	if app.Config.Dev {
		hostname = "localhost"
	}

	err = r.Run(fmt.Sprintf("%v:%v", hostname, app.Config.Port))
	if err != nil {
		panic(err)
	}
}
