export interface WeatherData {
	maxtemp_c: number
	mintemp_c: number
	avgtemp_c: number
	maxwind_kph: number
	avghumidity: number
	daily_chance_of_rain: number
}

export interface WeatherInfo {
	id: string
	weather: string
	windSpeed: number
}
