package server

import (
	"dashboard/config"
	"dashboard/message"
	"github.com/sirupsen/logrus"
	folderCreate "github.com/unjx-de/go-folder"
)

var Config = ServerConfig{}

const StorageDir = "storage/"
const IconsDir = StorageDir + "icons/"
const TemplatesDir = "templates/"

func Init() {
	createFolderStructure()
	config.ParseViperConfig(&Config, config.AddViperConfig("server"))
}

func createFolderStructure() {
	folders := []string{StorageDir, IconsDir, TemplatesDir}
	err := folderCreate.CreateFolders(folders, 0755)
	if err != nil {
		logrus.WithField("error", err).Fatal(message.CannotCreate.String())
	}
	logrus.WithField("folders", folders).Debug("folders created")
}
