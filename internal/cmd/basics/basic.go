package basics

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Version() {
	fmt.Println("workspace 1.0.0")
}

func Help() {
	fmt.Println("Help: This is a CLI application.")
	fmt.Println("Available Commands:")
	fmt.Printf("  %-10s - %s\n", "workspace", "Create workspace.")
	fmt.Printf("  %-10s - %s\n", "version", "Display the application version.")
	fmt.Printf("  %-10s - %s\n", "help", "Display help information.")
	fmt.Printf("  %-10s - %s\n", "clear", "Clear the console screen.")
	fmt.Printf("  %-10s - %s\n", "exit", "Exit the application.")
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
