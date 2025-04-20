package models

type Location struct {
    Name        string  `json:"name"`
    Region      string  `json:"region"`
    Country     string  `json:"country"`
    Latitude    float32 `json:"lat"`
    Longitude   float32 `json:"lon"`
    Timezone    string  `json:"tz_id"`
    LocalTime   string  `json:"localtime"`
}

type Current struct {
    LastUpdated     string      `json:"last_updated"`
    TempC           float64     `json:"temp_c"`
    TempF           float64     `json:"temp_f"`
    Condition       Condition   `json:"condition"`
    WindKph         float64     `json:"wind_kph"`
    WindDirection   string      `json:"wind_dir"`
    Humidity        int         `json:"humidity"`
    FeelsLikeC      float64     `json:"feelslike_c"`
    FeelsLikeF      float64     `json:"feelslike_f"`
}

type Condition struct {
    Text string `json:"text"`
    Icon string `json:"icon"`
}

type Forecast struct {
    ForecastDay     []ForeCastItem        `json:"forecastday"`
}

type ForeCastItem struct {
    Date        string          `json:"date"`
    Day         DayItem         `json:"day"`
}

type DayItem struct {
    MaxTempC            float64     `json:"maxtemp_c"`
    MinTempC            float64     `json:"mintemp_c"`
    AvgTempC            float64     `json:"avgtemp_c"`
    MaxWindKph          float64     `json:"maxwind_kph"`
    AvgHumidity         float64     `json:"avghumidity"`
    DailyChanceOfRain   float64     `json:"daily_chance_of_rain"`
    Condition           Condition   `json:"condition"`
}

type WeatherResponse struct {
    Location    Location    `json:"location"`
    Current     Current     `json:"current"`
    Forecast    Forecast    `json:"forecast"`
}