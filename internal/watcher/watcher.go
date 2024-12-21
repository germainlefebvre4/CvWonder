package watcher

import (
	"cvwonder/internal/cvparser"
	"cvwonder/internal/cvrender"
	"cvwonder/internal/utils"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
)

func ObserveFileEvents(outputDirectory string, inputFilePath string, themeName string, exportFormat string) {
	// setup watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	// Start the watcher
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// monitor only for write events
				if event.Op&fsnotify.Write == fsnotify.Write {
					logrus.Debug("Modification detected on template:", event.Name)
					content, err := cvparser.ParseFile(inputFilePath)
					utils.CheckError(err)

					cvrender.Render(content, outputDirectory, inputFilePath, themeName, exportFormat)
					utils.CheckError(err)
				}
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()

	// Add files to watch
	themeDirectory := fmt.Sprintf("themes/%s", themeName)
	err = watcher.AddWith(inputFilePath, fsnotify.WithBufferSize(65536))
	err = watcher.AddWith(themeDirectory+"/index.html", fsnotify.WithBufferSize(65536))
	<-done
}
