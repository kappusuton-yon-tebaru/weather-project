package weather

import (
	"github.com/Porsche-ths/weather-4-u/api-gateway/internal/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service,
	}
}

func (h *Handler) GetWeather(ctx *gin.Context) {
	reqParams := models.Coordinate{}

	if err := ctx.ShouldBindQuery(&reqParams); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	weather, err := h.service.GetWeather(reqParams)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, weather)
}
