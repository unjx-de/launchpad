package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetSessionCookie(c *gin.Context) {
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(sessionCookieName, createJWT(JwtExpiry), int(JwtExpiry.Seconds()), "/", strings.Split(c.Request.Host, ":")[0], true, true)
}

func DeleteSessionCookie(c *gin.Context) {
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(sessionCookieName, "none", -1, "/", strings.Split(c.Request.Host, ":")[0], true, true)
}
