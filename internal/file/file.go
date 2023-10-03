package file

import (
	"os"
	"path"
	"workspace/internal/config"
)

var filePath string
var file *os.File

var fullPathFile string

func Open(fileName string) {

	fullPathFile = path.Join(config.PathDirs["workspaces"], fileName)
	f, err := os.Create(fullPathFile)

	if err != nil {
		println(err)
		return
	}

	file = f
}

func Write(data []byte) {
	var content []byte

	if isFileEmpty() {
		content = data
	} else {
		content = append([]byte("\n\n"), data...)
	}

	_, err := file.Write(content)

	if err != nil {
		return
	}
}

func isFileEmpty() bool {
	fileInfo, err := file.Stat()

	if err != nil {
		return false
	}

	return fileInfo.Size() == 0
}

func Rename(oldName string, newName string) {
	oldNamePath := path.Join(config.PathDirs["workspaces"], oldName)
	newNamePath := path.Join(config.PathDirs["workspaces"], newName)

	err := os.Rename(oldNamePath, newNamePath)
	if err != nil {
		println("Error renaming the file:", err)
		return
	}
}

func Close() {
	file.Close()
}
