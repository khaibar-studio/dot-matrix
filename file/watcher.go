package file

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/khaibar-studio/dot-matrix/printer"
)

func Watch() {
	// Path to the directory to watch
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	pathSeparator := string(os.PathSeparator)
	directoryToWatch := home + pathSeparator + "Downloads" + pathSeparator

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
				execute(event.Name, directoryToWatch)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func execute(fileName, directoryToWatch string) {
	if !strings.Contains(fileName, "PRINT") {
		fmt.Println("OTHER FILE")
		return
	}
	printerName := strings.Replace(fileName, directoryToWatch+"PRINT_", "", -1)
	printerName = strings.Replace(printerName, "_", "\\", -1)
	printer.Exec(fileName, printerName)
	err := os.Remove(fileName)
	if err != nil {
		fmt.Printf("ERROR delete file: %s\n", fileName)
	}
}
