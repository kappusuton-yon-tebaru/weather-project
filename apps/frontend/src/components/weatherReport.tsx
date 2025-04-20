import { WeatherData } from "@/interfaces";

export default function WeatherReport({
  weatherJson,
}: {
  weatherJson: WeatherData;
}) {
  /*export interface WeatherData {
  maxtemp_c: number;  -- 
  mintemp_c: number; -- 
  avgtemp_c: number; -- 
  maxwind_kph: number; -- 
  avg_humidity: number;
  daily_chance_of_rain: number;
} */

  let rainChanceColor = "green-600";
  if (
    weatherJson.daily_chance_of_rain < 66 &&
    weatherJson.daily_chance_of_rain > 33
  ) {
    rainChanceColor = "amber-600";
  } else if (weatherJson.daily_chance_of_rain > 66) {
    rainChanceColor = "red-600";
  }
  return (
    <div
      className="flex flex-col items-center justify-center px-[56px] py-[28px] rounded-[24px] border-[4px] shadow-lg
     border-emerald-400 bg-emerald-50 text-[40px] "
    >
      <div className="flex flex-col space-y-[24px]">
        <div>
          <h1 className="text-green-600 w-full">Temperature</h1>

          <div className="flex flex-col px-[24px] w-full space-y-[24px]">
            <h1 className="text-emerald-950">
              Min: {weatherJson.mintemp_c} °C
            </h1>
            <h1 className="text-emerald-950">
              Max: {weatherJson.maxtemp_c} °C
            </h1>
            <h1 className="text-emerald-950">
              Average: {weatherJson.avgtemp_c} °C
            </h1>
          </div>
        </div>
        {/* <div className="flex flex-row space-x-[16px] w-full"> */}

        <div className=" flex flex-row space-x-[16px]">
          <h1 className="text-emerald-600">Max Wind Speed</h1>
          <h1 className="text-emerald-950">{weatherJson.maxwind_kph} kph.</h1>
        </div>
        <div className=" flex flex-row space-x-[16px]">
          <h1 className="text-emerald-600">Average Humidity</h1>
          <h1 className="text-emerald-950">{weatherJson.avghumidity} %</h1>
        </div>
      </div>
      <div
        className={`flex flex-row space-x-[16px] px-[56px] rounded-[24px] w-full ${
          rainChanceColor === "green-600"
            ? "bg-gradient-to-r from-transparent to-green-500"
            : rainChanceColor === "amber-600"
            ? "bg-gradient-to-r from-transparent to-amber-500"
            : "bg-gradient-to-r from-transparent to-red-500"
        }`}
      >
        <h1 className="text-emerald-600">Chance of Raining Today :</h1>
        <h1 className="text-white font-bold">
          {weatherJson.daily_chance_of_rain} %
        </h1>
      </div>
      {/* <div className="text-[48px]">Wind Speed: {weatherJson.windSpeed}</div> */}
    </div>
  );
}
