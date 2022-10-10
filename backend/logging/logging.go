package logging

import (
	"dashboard/config"
	"github.com/sirupsen/logrus"
)

var Config = LoggingConfig{}

const folder = "logging/"
const configFile = "logging.toml"

func Init() {
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006/01/02 15:04:05", FullTimestamp: true})
	config.AddViperConfig(folder, configFile)
	config.ParseViperConfig(&Config, configFile)
	setConfigLogLevel()
}

func setConfigLogLevel() {
	logLevel, err := logrus.ParseLevel(Config.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.FatalLevel)
	} else {
		logrus.SetLevel(logLevel)
	}
}
