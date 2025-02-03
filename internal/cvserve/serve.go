package cvserve

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"

	"github.com/germainlefebvre4/cvwonder/internal/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/jaschaephraim/lrserver"
	"github.com/sirupsen/logrus"
)

func (s *ServeServices) StartLiveReloader(port int, outputDirectory string, inputFilePath string) {
	// Default parameters
	if port == 0 {
		port = 8080
	}

	// Input file
	inputFilenameExt := path.Base(inputFilePath)
	inputFilename := inputFilenameExt[:len(inputFilenameExt)-len(path.Ext(inputFilenameExt))]

	// Create file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}
	defer watcher.Close()

	// Add dir to watcher
	outputFilename := filepath.Base(inputFilename) + ".html"
	err = watcher.AddWith(outputDirectory+"/"+outputFilename, fsnotify.WithBufferSize(65536*2))
	if err != nil {
		log.Fatalln(err)
	}

	// Create and start LiveReload server
	lr := lrserver.New(lrserver.DefaultName, lrserver.DefaultPort)
	go lr.ListenAndServe()

	if utils.CliArgs.Watch {
		// Start goroutine that requests reload upon watcher event
		go func() {
			for {
				select {
				case event := <-watcher.Events:
					if event.Op&fsnotify.Write == fsnotify.Write {
						lr.Reload(event.Name)
					}
				case err := <-watcher.Errors:
					log.Println(err)
				}
			}
		}()
	}

	// Start serving html
	s.StartServer(port, outputDirectory)
}

func (s *ServeServices) StartServer(port int, outputDirectory string) {
	if port == 0 {
		port = 8080
	}

	logrus.Debug(fmt.Sprintf("Listening on: http://localhost:%d", port))
	http.Handle("/", http.FileServer(http.Dir(outputDirectory)))
	listeningPort := fmt.Sprintf(":%d", port)
	http.ListenAndServe(listeningPort, nil)
}
