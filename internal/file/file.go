package file

import (
	"os"
)

var filePath string
var file *os.File
var fileName string = "dockerfile"

func Open() {
	f, err := os.Create(fileName)

	if err != nil {
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

func Close() {
	file.Close()
}

func SetFilePath(newFilePath string) {
	filePath = newFilePath
}

func resetVariables() {}
