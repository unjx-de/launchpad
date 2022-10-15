package auth

import (
	"dashboard/message"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CookieAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if noPasswordSet() || cookieAuthIsValid(c) {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, message.Response{Message: message.NotLoggedIn.String()})
			c.Abort()
		}
	}
}

func BlackListMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.GetHeader("X-Real-Ip")
		if ip == "" {
			ip = c.RemoteIP()
		}
		for _, v := range BlackList {
			if v.IP == ip && v.Amount >= maxLoginAttempts {
				c.JSON(http.StatusUnauthorized, message.Response{Message: message.NotLoggedIn.String()})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
