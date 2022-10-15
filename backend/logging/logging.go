package logging

import (
	"dashboard/config"
	"github.com/sirupsen/logrus"
)

var Config = LoggingConfig{}

func Init() {
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006/01/02 15:04:05", FullTimestamp: true})
	config.ParseViperConfig(&Config, config.AddViperConfig("logging"))
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
