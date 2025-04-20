package models

type Forecast struct {
	ForecastDay []ForeCastItem `json:"forecastday"`
}

type ForeCastItem struct {
	Date string  `json:"date"`
	Day  DayItem `json:"day"`
}

type DayItem struct {
	MaxTempC          float64 `json:"maxtemp_c"`
	MinTempC          float64 `json:"mintemp_c"`
	AvgTempC          float64 `json:"avgtemp_c"`
	MaxWindKph        float64 `json:"maxwind_kph"`
	AvgHumidity       float64 `json:"avghumidity"`
	DailyChanceOfRain float64 `json:"daily_chance_of_rain"`
}

type Weather struct {
	Forecast Forecast `json:"forecast"`
}

type WeatherResponse struct {
	Result Weather `json:"result"`
}
