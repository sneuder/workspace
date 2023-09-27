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

	fullPathFile = path.Join(config.BasePath, config.Dirs["workspace"], config.Dirs["workspaces"], fileName)
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

func Close() {
	file.Close()
}
