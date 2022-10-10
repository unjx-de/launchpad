package main

import (
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
		if logrus.GetLevel() != logrus.TraceLevel {
			return
		}
		reqUri := c.Request.RequestURI
		if strings.Contains(reqUri, "/storage") {
			return
		}
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		logrus.WithFields(logrus.Fields{
			"status":  http.StatusText(c.Writer.Status()),
			"latency": latencyTime,
			"client":  c.RemoteIP(),
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
		AllowHeaders:     []string{"content-type"},
		AllowMethods:     []string{"GET"},
	}))

	router.Use(myLogger())
	_ = router.SetTrustedProxies(server.Config.AllowedHosts)
	logrus.WithField("allowedOrigins", server.Config.AllowedHosts).Debug("middlewares set")
}
