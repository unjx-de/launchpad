import { bookmark_Bookmark, system_LiveInformation, system_StaticInformation, weather_OpenWeatherApiResponse } from "~/services/openapi";

export const emptyWeather: weather_OpenWeatherApiResponse = {
  main: { humidity: 0, temp: 0 },
  name: "",
  sys: { sunrise: 0, sunset: 0 },
  units: "°C",
  weather: [{ description: "", icon: "01d" }],
};

export const emptyLiveSystem: system_LiveInformation = {
  cpu: { percentage: [], value: "" },
  disk: { percentage: 0, value: "" },
  ram: { percentage: 0, value: "" },
  server_uptime: 0,
};

export const emptyStaticSystem: system_StaticInformation = {
  cpu: { architecture: "", name: "", threads: 0 },
  disk: { readable: "", unit: 0, unit_string: "", value: 0 },
  ram: { readable: "", unit: 0, unit_string: "", value: 0 },
};

export const emptyBookmark: bookmark_Bookmark = { icon: "", name: "", url: "" };

export const PercentageArraySize = 120;
