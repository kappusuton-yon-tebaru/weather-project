package router

import (
	"github.com/Porsche-ths/weather-4-u/weather-service/cmd/weather/app"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func New() *Router {
	return &Router{
		gin.Default(),
	}
}

func (r *Router) RegisterRoutes(app *app.App) {
	api := r.Group("/api")
	api.GET("/weather", app.WeatherHandler.GetWeather)
	// api.GET("/ip", app.GeoipHandler.CoordinateFromIP)
}
