package geocode

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

func (h *Handler) SearchCoordinate(ctx *gin.Context) {
	query := ctx.Query("text")

	locations, err := h.svc.SearchFromAddress(ctx, query)
	if err != nil {
		werr := werror.Must(err)
		ctx.JSON(werr.Code(500), werr)
		return
	}

	ctx.JSON(200, map[string]any{
		"results": locations,
	})
}
