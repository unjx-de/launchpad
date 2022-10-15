package auth

type AuthConfig struct {
	Auth Auth
}

type Auth struct {
	Password string
	Secret   string
}
