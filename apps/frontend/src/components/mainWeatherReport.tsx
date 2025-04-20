"use client";

// import Image from "next/image";
import { getWeatherInfo } from "@/libs/getWeatherInfo";
import { useState } from "react";
import WeatherReport from "@/components/weatherReport";
import { WeatherInfo, WeatherData } from "@/interfaces";
import axios from "axios";
import { isValidWeatherInfo } from "@/utils/validators";

export default function MainWeatherReport() {
  const [weatherData, setWeatherData] = useState<WeatherData | null>(null);
  const [address, setAddress] = useState("");

  const handleGetWeather = async () => {
    try {
      // Fetch IP address
      const ipResponse = await axios.get("https://api.ipify.org?format=json");
      const ipAddress = ipResponse.data.ip;

      // Fetch weather data using the IP address

      const data = await getWeatherInfo({ address: address, IP: ipAddress });
      // if (isValidWeatherInfo(data)) {
      //   setWeatherData(data);
      // } else {
      //   setWeatherData(null); // or handle it as needed
      // }

      // const data = await getWeatherInfo({ address: "test", IP: "test" });
      setWeatherData(data);
    } catch (error) {
      console.error("Failed to fetch IP or weather data:", error);
    }
  };

  return (
    <div className=" flex flex-col justify-center space-y-[32px] bg-green-200 px-[32px] py-[24px] rounded-[24px]">
      {/* <div>Get Weather Report By clicking the button</div> */}
      <div className="w-auto flex justify-center flex-col space-y-[32px]">
        <div className="flex flex-col space-y-[32px] justify-center">
          <label
            htmlFor="addressField"
            className="text-[32px] font-bold text-emerald-950"
          >
            Insert Your City Name or Your Province
          </label>
          <input
            type="text"
            id="addressField"
            value={address}
            placeholder="Type your city name or your province"
            className="px-[32px] bg-white rounded-[8px] py-[16px] border-2 text-green-800 placeholder:text-gray-400
             border-green-600 focus:ring-0 focus:border-green-800 outline-none focus:outline-none"
            onChange={(e) => setAddress(e.target.value)}
          />
        </div>

        <button
          onClick={handleGetWeather}
          className=" bg-green-600 text-[24px] font-bold px-[24px] py-[16px] rounded-[8px] text-green-50  hover:bg-green-700  active:bg-green-800"
        >
          {weatherData ? "Get New Weather" : "Get Weather"}
        </button>
      </div>

      {weatherData ? (
        <div>
          <WeatherReport weatherJson={weatherData} />
        </div>
      ) : (
        ""
      )}
    </div>
  );
}
