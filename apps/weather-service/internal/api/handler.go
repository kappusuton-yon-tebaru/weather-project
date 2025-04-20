package api

import (
	// "log"
	// "net/http"

	"github.com/Porsche-ths/weather-4-u/weather-service/internal/werror"
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

// func (h *Handler) SearchCoordinate(ctx *gin.Context) {
// 	query := ctx.Query("text")

// 	locations, err := h.svc.SearchFromAddress(ctx, query)
// 	if err != nil {
// 		werr := werror.Must(err)
// 		ctx.JSON(werr.Code(500), werr)
// 		return
// 	}

// 	ctx.JSON(200, map[string]any{
// 		"results": locations,
// 	})
// }

func (h *Handler) GetWeather(ctx *gin.Context) {
	// query := ctx.Query("text")

	// locations, err := h.svc.SearchFromAddress(ctx, query)
	// if err != nil {
	// 	werr := werror.Must(err)
	// 	ctx.JSON(werr.Code(500), werr)
	// 	return
	// }

	// ctx.JSON(200, map[string]any{
	// 	"results": locations,
	// })
	///////////////////////
	// Build the API request URL
	query := ctx.Query("q")
	days := ctx.Query("days")

	data, err := h.svc.FetchWeather(ctx, query, days)
	if err != nil {
		werr := werror.Must(err)
		ctx.JSON(werr.Code(500), werr)
		return
	}

	ctx.JSON(200, map[string]any{
		"result": data,
	})

    // url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

    // // Make the request
    // resp, err := http.Get(url)
    // if err != nil {
    //     log.Fatalf("Failed to make request: %v", err)
    // }
    // defer resp.Body.Close()

    // // Check if the request was successful
    // if resp.StatusCode != http.StatusOK {
    //     log.Fatalf("Unexpected response status: %s", resp.Status)
    // }

    // // Decode the JSON response
    // var weather models.WeatherResponse
    // if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
    //     log.Fatalf("Failed to decode JSON: %v", err)
    // }
}
