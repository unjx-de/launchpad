package weather

type WeatherConfig struct {
	Location    Location
	OpenWeather OpenWeather `mapstructure:"OPEN_WEATHER"`
}

type Location struct {
	Latitude  float32
	Longitude float32
}

type OpenWeather struct {
	Key   string
	Units string
	Lang  string
}

type OpenWeatherApiResponse struct {
	Weather []OpenWeatherApiWeather `json:"weather" validate:"required"`
	Main    OpenWeatherApiMain      `json:"main" validate:"required"`
	Sys     OpenWeatherApiSys       `json:"sys" validate:"required"`
	Name    string                  `json:"name" validate:"required"`
	Units   string                  `json:"units" validate:"required"`
}

type OpenWeatherApiWeather struct {
	Description string `json:"description" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
}

type OpenWeatherApiMain struct {
	Temp     float32 `json:"temp" validate:"required"`
	Humidity float32 `json:"humidity" validate:"required"`
}

type OpenWeatherApiSys struct {
	Sunrise uint64 `json:"sunrise" validate:"required"`
	Sunset  uint64 `json:"sunset" validate:"required"`
}
