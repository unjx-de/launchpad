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
