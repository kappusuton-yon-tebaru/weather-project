package geolocation

import (
	"github.com/Porsche-ths/weather-4-u/api-gateway/internal/api/weather"
	"github.com/Porsche-ths/weather-4-u/api-gateway/internal/models"
	"github.com/Porsche-ths/weather-4-u/api-gateway/internal/werror"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	geoService *Service
	weaService *weather.Service
}

func NewHandler(geoService *Service, weaService *weather.Service) *Handler {
	return &Handler{
		geoService,
		weaService,
	}
}

func (h *Handler) GetLocation(ctx *gin.Context) {
	reqParams := models.RequestParams{
		IpAddress: ctx.Query("ip_address"),
		Address:   ctx.Query("address"),
	}

	location, err := h.geoService.GetLocation(reqParams)
	if err != nil {
		werror := werror.Must(err)
		ctx.JSON(werror.Code(500), werror)
		return
	}

	if location == (models.Geolocation{}) {
		ctx.JSON(200, "No location found")
		return
	}

	weather, err := h.weaService.GetWeather(location.Coordinate)
	if err != nil {
		werror := werror.Must(err)
		ctx.JSON(werror.Code(500), werror)
		return
	}

	forcastDayData := weather.Result.Forecast.ForecastDay
	if len(forcastDayData) == 0 {
		ctx.JSON(200, "No forecast data")
		return
	}

	ctx.JSON(200, forcastDayData[0].Day)
}
