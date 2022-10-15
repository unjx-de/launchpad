package main

import (
	"dashboard/auth"
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
		authGroup := api.Group("/auth")
		{
			authGroup.Use(auth.BlackListMiddleware())
			authGroup.POST("login", login)
			authGroup.POST("logout", logout)
		}
		bookmarkGroup := api.Group("/bookmarks")
		{
			bookmarkGroup.Use(auth.CookieAuthRequired())
			bookmarkGroup.GET("", getBookmarks)
		}
		weatherGroup := api.Group("/weather")
		{
			weatherGroup.Use(auth.CookieAuthRequired())
			weatherGroup.GET("", getWeather)
		}
		systemGroup := api.Group("/system")
		{
			systemGroup.Use(auth.CookieAuthRequired())
			systemGroup.GET("/static", routeStaticSystem)
			systemGroup.GET("/live", routeLiveSystem)
			systemGroup.GET("/ws", serveWs)
		}
	}
	staticGroup := router.Group("/static")
	{
		staticGroup.Use(auth.CookieAuthRequired())
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
