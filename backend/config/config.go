package config

import (
	"dashboard/message"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

func AddViperConfig(name string) string {
	file := name + ".json"
	viper.SetConfigFile(name + "/" + file)
	err := viper.MergeInConfig()
	if err != nil {
		logrus.WithField("file", name+".json").Fatal(message.CannotOpen.String())
	}
	return file
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
