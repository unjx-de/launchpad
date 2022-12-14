package main

import (
	"dashboard/auth"
	"dashboard/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func myLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqUri := c.Request.RequestURI
		if strings.Contains(reqUri, "/static") {
			return
		}
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		logrus.WithFields(logrus.Fields{
			"status":  http.StatusText(c.Writer.Status()),
			"latency": latencyTime,
			"client":  auth.GetRealIp(c),
			"method":  c.Request.Method,
		}).Trace(reqUri)
	}
}

func cacheControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	}
}

func setMiddlewares(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     server.Config.AllowedHosts,
		AllowCredentials: true,
		AllowHeaders:     []string{"content-type", "password"},
		AllowMethods:     []string{"GET", "POST"},
	}))

	if logrus.GetLevel() == logrus.TraceLevel {
		router.Use(myLogger())
	}
	_ = router.SetTrustedProxies(server.Config.AllowedHosts)
	logrus.WithField("allowedOrigins", server.Config.AllowedHosts).Debug("middlewares set")
}
