package bookmark

import (
	"dashboard/message"
	"dashboard/server"
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Bookmarks []Bookmark

const folder = "bookmark/"
const desiredFolder = server.StorageDir
const bookmarksFile = "bookmarks.json"

func Init() {
	parseBookmarks()
	go watchBookmarks()
}

func copyBookmarks() {
	source, _ := os.Open(folder + bookmarksFile)
	defer source.Close()
	destination, err := os.Create(desiredFolder + bookmarksFile)
	if err != nil {
		logrus.WithField("file", bookmarksFile).Error(message.CannotCreate.String())
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		logrus.WithField("file", bookmarksFile).Error(message.CannotCreate.String())
	}
}

func parseBookmarks() {
	jsonFile, err := os.Open(desiredFolder + bookmarksFile)
	if err != nil {
		copyBookmarks()
		jsonFile, err = os.Open(desiredFolder + bookmarksFile)
		if err != nil {
			logrus.WithField("file", bookmarksFile).Error(message.CannotOpen.String())
			return
		}
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		logrus.WithField("file", bookmarksFile).Error(message.CannotProcess.String())
		return
	}
	err = json.Unmarshal(byteValue, &Bookmarks)
	if err != nil {
		logrus.WithField("file", bookmarksFile).Error(message.CannotProcess.String())
		return
	}
}

func watchBookmarks() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logrus.WithField("watcher", err).Fatal(message.CannotCreate.String())
	}
	defer watcher.Close()
	done := make(chan bool)

	go func() {
		for {
			select {
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logrus.WithField("error", err).Error(message.CannotProcess.String())
			case _, ok := <-watcher.Events:
				if !ok {
					return
				}
				parseBookmarks()
				logrus.WithField("bookmarks", len(Bookmarks)).Trace(bookmarksFile + " changed")
			}
		}
	}()

	if err := watcher.Add(desiredFolder + bookmarksFile); err != nil {
		logrus.WithField("watcher", err).Fatal(message.CannotCreate.String())
	}
	<-done
}
