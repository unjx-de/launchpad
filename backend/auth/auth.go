package auth

import (
	"crypto/rand"
	"dashboard/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

var Config = AuthConfig{}
var BlackList []Client

const (
	JwtExpiry         = 30 * 24 * time.Hour
	sessionCookieName = "launchpad-session"
	secretLength      = 40
	maxLoginAttempts  = 3
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

func ResetBlackList(ip string) {
	for i, v := range BlackList {
		if v.IP == ip {
			BlackList[i].Amount = 1
		}
		break
	}
}

func CheckBlackList(ip string) {
	for i, v := range BlackList {
		if v.IP == ip {
			if v.Amount < maxLoginAttempts {
				BlackList[i].Amount++
			}
			return
		}
	}
	BlackList = append(BlackList, Client{IP: ip, Amount: 1})
}

func GetRealIp(c *gin.Context) string {
	ip := c.GetHeader("X-Real-Ip")
	if ip == "" {
		ip = c.RemoteIP()
	}
	return ip
}
