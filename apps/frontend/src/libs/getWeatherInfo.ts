import { WeatherData } from "@/interfaces";
import { env } from "next-runtime-env";

export async function getWeatherInfo({
  address,
  IP,
}: {
  address?: string;
  IP: string;
}) {
  const gatewayEndpoint = env("NEXT_PUBLIC_GATEWAY_ENDPOINT");
  console.log("ip is " + IP + " address is " + address);
  console.log(gatewayEndpoint);

  const mockReturnData: WeatherData = {
    mintemp_c: 10,
    maxtemp_c: 10,
    avgtemp_c: 10,
    maxwind_kph: 10,
    avghumidity: 10,
    daily_chance_of_rain: 90,
  };
  if (address === "test" || IP === "test") {
    return mockReturnData;
  }
  if (!address) {
    const response = await fetch(
      `${gatewayEndpoint}/location?ip_address=${IP}`,
    );
    if (!response) {
      throw new Error("Failed to fetch weather data");
    }
    return response.json();
  } else {
    const response = await fetch(
      `${gatewayEndpoint}/location?address=${address}`,
      {
        method: "GET", // Ensure this matches the method you're using in the browser
        headers: {
          "Content-Type": "application/json", // Include any necessary headers
        },
      },
    );

    if (!response) {
      throw new Error("Failed to fetch weather data");
    }
    return response.json();
  }
}
