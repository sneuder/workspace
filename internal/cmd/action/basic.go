package action

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Version() {
	fmt.Println("1.0.0")
}

func Help() {
	fmt.Println("Help: This is a CLI application.")
	fmt.Println("Available Commands:")
	fmt.Println("  version - Display the application version.")
	fmt.Println("  help    - Display help information.")
	fmt.Println("  clear   - Clear the console screen.")
	fmt.Println("  exit    - Exit the application.")
}

func Clear() {
	cmd := exec.Command("clear")

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Exit() {
	fmt.Println("Exiting the workspace.")
	os.Exit(0)
}
