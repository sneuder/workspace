package helpers

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"workspace/constants"
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

func GetWkspaceFolder() string {
	folderDir := GetHomeDir()
	return filepath.Join(folderDir, constants.FOLDER_WKSPACE)
}
