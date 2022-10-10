package main

import (
	"dashboard/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

func handleStaticFiles(router *gin.Engine) {
	router.LoadHTMLFiles(server.TemplatesDir + "index.html")
	router.StaticFile("robots.txt", "templates/robots.txt")
	router.StaticFile("favicon.ico", "templates/favicon.ico")
	serveFoldersInTemplates(router)
}

func serveFoldersInTemplates(router *gin.Engine) {
	static := router.Group("")
	_ = filepath.WalkDir(server.TemplatesDir, func(path string, info os.DirEntry, err error) error {
		if info.IsDir() && info.Name() != strings.TrimSuffix(server.TemplatesDir, "/") {
			static.Use(cacheControl())
			static.Static(info.Name(), server.TemplatesDir+info.Name())
			logrus.WithField("folder", info.Name()).Debug("serve static folder")
		}
		return err
	})
}
