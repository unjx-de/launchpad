package weather

import (
	"dashboard/config"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

var Config = WeatherConfig{}
var CurrentOpenWeather = OpenWeatherApiResponse{}

func Init() {
	config.ParseViperConfig(&Config, config.AddViperConfig("weather"))
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
			logrus.Error("weather cannot be updated")
		} else {
			body, _ := io.ReadAll(resp.Body)
			err = json.Unmarshal(body, &CurrentOpenWeather)
			if err != nil {
				logrus.Error("weather cannot be processed")
			} else {
				logrus.WithFields(logrus.Fields{"temp": fmt.Sprintf("%0.2f%s", CurrentOpenWeather.Main.Temp, CurrentOpenWeather.Units), "location": CurrentOpenWeather.Name}).Trace("weather updated")
			}
			resp.Body.Close()
		}
		time.Sleep(interval)
	}
}
