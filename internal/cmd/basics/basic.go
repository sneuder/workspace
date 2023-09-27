package basics

import (
	"os"
	"os/exec"
	"runtime"
)

func Version() {
	println("1.0.0")
}

func Help() {
	println("Help: This is a CLI application.")
	println("Available Commands:")
	println("  workspace - Create workspace.")
	println("  version - Display the application version.")
	println("  help    - Display help information.")
	println("  clear   - Clear the console screen.")
	println("  exit    - Exit the application.")
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
	println("Exiting the workspace.")
	os.Exit(0)
}
