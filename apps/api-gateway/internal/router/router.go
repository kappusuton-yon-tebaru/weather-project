package router

import (
	"github.com/Porsche-ths/weather-4-u/api-gateway/cmd/app"
	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}))

	api := r.Group("/api")
	api.GET("/location", app.GeolocationHandler.GetLocation)
	api.GET("/weather", app.WeatherHandler.GetWeather)
}
