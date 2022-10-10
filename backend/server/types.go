package server

type ServerConfig struct {
	Port         int
	AllowedHosts []string `mapstructure:"ALLOWED_HOSTS"`
	Swagger      bool
}
