package main

import (
	"dashboard/docs"
	"dashboard/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"net/url"
)

func setupSwagger(router *gin.Engine) {
	if server.Config.Swagger {
		docs.SwaggerInfo.Title = "Launchpad"
		docs.SwaggerInfo.Version = "1.0.2"
		docs.SwaggerInfo.BasePath = "/api"
		parsed, _ := url.Parse(server.Config.AllowedHosts[0])
		docs.SwaggerInfo.Host = parsed.Host

		// @contact.url   https://github.com/unjx-de/launchpad
		// @license.name  MIT License
		// @license.url   https://github.com/unjx-de/launchpad/blob/main/LICENSE

		router.GET("/swagger", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		})
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		logrus.WithField("url", server.Config.AllowedHosts[0]+"/swagger").Debug("swagger running")
	}
}
