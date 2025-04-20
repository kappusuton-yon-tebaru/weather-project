package router

import (
	"github.com/Porsche-ths/weather-4-u/geolocation-service/cmd/geolocation/app"
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
	api.GET("/search", app.GeocodeHandler.SearchCoordinate)
	api.GET("/ip", app.GeoipHandler.CoordinateFromIP)
}
