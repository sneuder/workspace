package helpers

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return homeDir
}

func CreateFolder(pathFolder string, folerName string) {
	fullPathNameFolder := filepath.Join(pathFolder, folerName)
	exec.Command("mkdir", "-p", fullPathNameFolder).Run()
}
