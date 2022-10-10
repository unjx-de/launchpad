package main

import (
	"dashboard/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func setupRouter() *gin.Engine {
	router := gin.New()
	setMiddlewares(router)
	handleStaticFiles(router)

	api := router.Group("/api")
	{
		bookmarkGroup := api.Group("/bookmarks")
		{
			bookmarkGroup.GET("", getBookmarks)
		}
		weatherGroup := api.Group("/weather")
		{
			weatherGroup.GET("", getWeather)
		}
		systemGroup := api.Group("/system")
		{
			systemGroup.GET("/static", routeStaticSystem)
			systemGroup.GET("/live", routeLiveSystem)
			systemGroup.GET("/ws", serveWs)
		}
	}
	staticGroup := router.Group("/static")
	{
		staticGroup.Use(cacheControl())
		staticGroup.Static("", server.IconsDir)
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/")
	})
	logrus.WithField("amount", len(router.Routes())).Debug("routes initialized")
	return router
}
