package main

import (
	"dashboard/bookmark"
	"dashboard/logging"
	"dashboard/message"
	"dashboard/server"
	"dashboard/system"
	"dashboard/weather"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logging.Init()
	server.Init()
	weather.Init()
	bookmark.Init()
	system.Init()

	gin.SetMode(gin.ReleaseMode)
	router := setupRouter()
	setupSwagger(router)

	logrus.WithField("port", server.Config.Port).Info("application running")
	err := router.Run(fmt.Sprintf(":%d", server.Config.Port))
	if err != nil {
		logrus.WithField("error", err).Fatal(message.CannotStart.String())
	}
}
