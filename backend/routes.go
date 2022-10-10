package main

import (
	"dashboard/bookmark"
	"dashboard/message"
	"dashboard/system"
	"dashboard/weather"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
)

// @Schemes
// @Summary     get all bookmarks
// @Description gets all bookmarks as array
// @Tags        bookmarks
// @Produce     json
// @Success     200 {array}  bookmark.Bookmark
// @Success     204 {object} message.Response
// @Router      /bookmarks [get]
func getBookmarks(c *gin.Context) {
	c.JSON(http.StatusOK, bookmark.Bookmarks)
}

// @Schemes
// @Summary     get the current weather
// @Description gets the current weather
// @Tags        weather
// @Produce     json
// @Success     200 {array}  weather.OpenWeatherApiResponse
// @Success     204 {object} message.Response
// @Router      /weather [get]
func getWeather(c *gin.Context) {
	if weather.Config.OpenWeather.Key != "" {
		c.JSON(http.StatusOK, weather.CurrentOpenWeather)
	} else {
		c.JSON(http.StatusNoContent, message.Response{Message: message.CannotFind.String()})
	}
}

// @Schemes
// @Summary     static system information
// @Description gets static information of the system
// @Tags        system
// @Produce     json
// @Success     200 {object} system.StaticInformation
// @Router      /system/static [get]
func routeStaticSystem(c *gin.Context) {
	if system.Config.LiveSystem {
		c.JSON(http.StatusOK, system.Live.System.Static)
	} else {
		c.JSON(http.StatusNoContent, message.Response{Message: message.CannotFind.String()})
	}
}

// @Schemes
// @Summary     live system information
// @Description gets live information of the system
// @Tags        system
// @Produce     json
// @Success     200 {object} system.LiveInformation
// @Router      /system/live [get]
func routeLiveSystem(c *gin.Context) {
	if system.Config.LiveSystem {
		c.JSON(http.StatusOK, system.Live.System.Live)
	} else {
		c.JSON(http.StatusNoContent, message.Response{Message: message.CannotFind.String()})
	}
}

func serveWs(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.WithField("error", err).Warning("Cannot upgrade websocket")
		return
	}
	messageChan := make(system.NotifierChan)
	system.Live.Hub.NewClients <- messageChan
	defer func() {
		system.Live.Hub.ClosingClients <- messageChan
		conn.Close()
	}()
	go readPump(conn)
	for {
		select {
		case msg, ok := <-messageChan:
			if !ok {
				err := conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					return
				}
				return
			}
			err := conn.WriteJSON(msg)
			if err != nil {
				return
			}
		}
	}
}
