package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: print <printerName>")
		return
	}

	printerName := os.Args[1]

	// Sample text with escape sequences
	data := "\x1B@Hello, Dot Matrix Printer!\n" // ESC @ (initialize) and text
	data += "\x1BEBold Text\x1BF\n"             // ESC E (bold on), text, and ESC F (bold off)
	data += "Normal Text\n"

	// Create a temporary file to store the print job
	tempFile, err := os.CreateTemp("", "print_job_*.txt")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.Write([]byte(data))
	if err != nil {
		fmt.Println("Error writing to temporary file:", err)
		return
	}
	tempFile.Close()

	// Define the command based on the OS
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", fmt.Sprintf("type %s > %s", tempFile.Name(), printerName))
	case "darwin", "linux":
		cmd = exec.Command("lp", "-d", printerName, tempFile.Name())
	default:
		fmt.Println("Unsupported platform")
		return
	}

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error printing:", err)
	} else {
		fmt.Println("Print job sent successfully!")
	}
}
