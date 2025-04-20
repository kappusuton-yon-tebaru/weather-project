package main

import (
	"fmt"

	"github.com/Porsche-ths/weather-4-u/geolocation-service/cmd/geolocation/app"
	"github.com/Porsche-ths/weather-4-u/geolocation-service/internal/router"
)

func main() {
	app, err := app.Intialize()
	if err != nil {
		panic(err)
	}

	fmt.Println(app.Config)

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
