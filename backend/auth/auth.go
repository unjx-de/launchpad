package auth

import (
	"crypto/rand"
	"dashboard/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

var Config = AuthConfig{}

const (
	JwtExpiry         = 30 * 24 * time.Hour
	sessionCookieName = "launchpad-session"
	secretLength      = 40
)

func Init() {
	config.ParseViperConfig(&Config, config.AddViperConfig("auth"))
	if noSecretSet() {
		Config.Auth.Secret = secretGenerator()
	}
}

func secretGenerator() string {
	logrus.Warning("no secret set, generating...")
	b := make([]byte, secretLength)
	_, _ = rand.Read(b)
	return string(b)
}

func noSecretSet() bool {
	return len(Config.Auth.Secret) == 0
}

func noPasswordSet() bool {
	return len(Config.Auth.Password) == 0
}

func cookieAuthIsValid(c *gin.Context) bool {
	token, err := c.Cookie(sessionCookieName)
	if err != nil {
		return false
	}
	return validateJWT(token) == nil
}
