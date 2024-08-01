package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// Path to the directory to watch
	directoryToWatch := "./queue"

	// Create a new watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("error creating watcher: %v", err)
	}
	defer watcher.Close()

	// Add the directory to the watcher
	err = watcher.Add(directoryToWatch)
	if err != nil {
		log.Fatalf("error adding directory to watcher: %v", err)
	}

	fmt.Printf("Watching directory: %s\n", directoryToWatch)

	// Start watching for events
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				fmt.Printf("New file created: %s\n", event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Printf("Error: %v\n", err)
		}
	}
}
