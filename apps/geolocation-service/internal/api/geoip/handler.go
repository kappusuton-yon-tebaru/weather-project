package geoip

import (
	"github.com/Porsche-ths/weather-4-u/geolocation-service/internal/werror"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{
		svc,
	}
}

func (h *Handler) CoordinateFromIP(ctx *gin.Context) {
	ip := ctx.Query("ip")

	location, err := h.svc.CoordinateFromIP(ctx, ip)
	if err != nil {
		werr := werror.Must(err)
		ctx.JSON(werr.Code(500), werr)
		return
	}

	ctx.JSON(200, map[string]any{
		"result": location,
	})
}
