package action

import (
	"fmt"
	"workspace/internal/file"
)

var instructions []string

func Workspace() {
	fmt.Println("Creating workspace")

	setInstructions()

	file.Open()

	for _, value := range instructions {
		fmt.Println(value)
	}

	data := []byte("This is a sample byte slice.")
	file.Write(data)
}

func setInstructions() {
	instructions = append(instructions, "FROM ubuntu:latest")
	instructions = append(instructions, `CMD ["sleep", "infinity"]`)
}

func addInstructions() {

}
