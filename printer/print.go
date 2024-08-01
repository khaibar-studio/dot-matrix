package printer

import (
	"fmt"
	"os/exec"
	"runtime"
)

func Exec(fileName, printerName string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", fmt.Sprintf("type %s > %s", fileName, printerName))
	case "darwin", "linux":
		cmd = exec.Command("lp", "-d", printerName, fileName)
	default:
		fmt.Println("Unsupported platform")
		return
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error printing:", err)
	} else {
		fmt.Println("Print job sent successfully!")
	}
}
