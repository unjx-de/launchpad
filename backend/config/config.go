package config

import (
	"dashboard/message"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

func AddViperConfig(folder string, configFile string) {
	viper.SetConfigFile(folder + configFile)
	err := viper.MergeInConfig()
	if err != nil {
		logrus.WithField("file", configFile).Fatal(message.CannotOpen.String())
	}
}

func ParseViperConfig(config interface{}, configFile string) {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err := viper.Unmarshal(config, viper.DecodeHook(mapstructure.StringToSliceHookFunc(",")))
	if err != nil {
		logrus.WithField("file", configFile).Fatal(message.CannotProcess.String())
	}
	logrus.WithField("file", configFile).Debug("config file parsed")
}
