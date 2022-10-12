package weather

import (
	"dashboard/config"
	"dashboard/message"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

var Config = WeatherConfig{}
var CurrentOpenWeather = OpenWeatherApiResponse{}

const folder = "weather/"
const configFile = "weather.toml"

func Init() {
	config.AddViperConfig(folder, configFile)
	config.ParseViperConfig(&Config, configFile)
	if Config.OpenWeather.Key != "" {
		setWeatherUnits()
		go updateWeather(time.Second * 150)
	}
}

func setWeatherUnits() {
	if Config.OpenWeather.Units == "imperial" {
		CurrentOpenWeather.Units = "°F"
	} else {
		CurrentOpenWeather.Units = "°C"
	}
}

func updateWeather(interval time.Duration) {
	for {
		resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=%s&lang=%s",
			Config.Location.Latitude,
			Config.Location.Longitude,
			Config.OpenWeather.Key,
			Config.OpenWeather.Units,
			Config.OpenWeather.Lang))
		if err != nil {
			logrus.WithField("error", err).Error(message.CannotGet.String())
		} else {
			body, _ := io.ReadAll(resp.Body)
			err = json.Unmarshal(body, &CurrentOpenWeather)
			if err != nil {
				logrus.WithField("error", err).Error(message.CannotProcess.String())
			} else {
				logrus.WithField("temp", fmt.Sprintf("%0.2f%s", CurrentOpenWeather.Main.Temp, CurrentOpenWeather.Units)).Trace("weather updated")
			}
		}
		resp.Body.Close()
		time.Sleep(interval)
	}
}
