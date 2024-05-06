package file

import (
	"fmt"
	"os"
	"path"
	"workspace/internal/config"
)

var fullPathFile string

func OpenFile(fileName string, pathDir string) (error, *os.File) {
	fullPathFile = path.Join(pathDir, fileName)
	f, err := os.Create(fullPathFile)

	if err != nil {
		return err, f
	}

	return nil, f
}

func WriteFile(data []byte, file *os.File) {
	var content []byte

	if isEmptyFile(file) {
		content = data
	} else {
		content = append([]byte("\n"), data...)
	}

	_, err := file.Write(content)

	if err != nil {
		return
	}
}

func ReadFile(filePath string) string {
	data, _ := os.ReadFile(filePath)
	return string(data)
}

func isEmptyFile(file *os.File) bool {
	fileInfo, err := file.Stat()

	if err != nil {
		return false
	}

	return fileInfo.Size() == 0
}

func RenameFile(oldName string, newName string) {
	oldNamePath := path.Join(config.PathDirs["workspaces"], oldName)
	newNamePath := path.Join(config.PathDirs["workspaces"], newName)

	err := os.Rename(oldNamePath, newNamePath)
	if err != nil {
		fmt.Println("Error renaming the file:", err)
	}
}

func CloseFile(file *os.File) {
	file.Close()
}

func ExistsFile(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func RemoveFile(fileName string) {
	err := os.Remove(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
