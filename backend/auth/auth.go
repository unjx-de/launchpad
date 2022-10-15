package auth

import (
	"dashboard/config"
)

var Config = AuthConfig{}

func Init() {
	config.ParseViperConfig(&Config, config.AddViperConfig("auth"))
}
