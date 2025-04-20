import { WeatherInfo, WeatherData } from "@/interfaces";
/*WeatherData {
  maxtemp_c: number;
  mintemp_c: number;
  avgtemp_c: number;
  maxwind_kph: number;
  avg_humidity: number;
  daily_chance_of_rain: number;
} */
export function isValidWeatherInfo(data: WeatherData): boolean {
  return (
    data !== null &&
    data.maxtemp_c !== null &&
    data.mintemp_c !== null &&
    data.avgtemp_c !== null &&
    data.maxwind_kph !== null &&
    data.avghumidity !== null &&
    data.daily_chance_of_rain !== null &&
    typeof data.maxtemp_c === "number" &&
    typeof data.mintemp_c === "number" &&
    typeof data.avgtemp_c === "number" &&
    typeof data.maxwind_kph === "number" &&
    typeof data.avghumidity === "number" &&
    typeof data.daily_chance_of_rain === "number"
  );
}
