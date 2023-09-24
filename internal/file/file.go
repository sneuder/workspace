package file

import (
	"fmt"
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
	n, err := file.Write(data)

	if err != nil {
		return
	}
	fmt.Printf("n: %v\n", n)
}

func Close() {

}

func SetFilePath(newFilePath string) {
	filePath = newFilePath
}

func resetVariables() {}
