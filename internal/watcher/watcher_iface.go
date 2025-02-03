package watcher

import "github.com/germainlefebvre4/cvwonder/internal/cvrender"

type WatcherInterface interface {
	ObserveFileEvents(renderService cvrender.RenderInterface, baseDirectory string, outputDirectory string, inputFilePath string, themeName string, exportFormat string)
}

type WatcherServices struct{}

func NewWatcherServices() (WatcherInterface, error) {
	return &WatcherServices{}, nil
}
