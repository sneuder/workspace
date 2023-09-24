package action

import (
	"fmt"
	"workspace/internal/docker"
)

// var instructions []string

func Workspace() {
	fmt.Println("Creating workspace")

	docker.StartImageProcess()
	docker.StartContainerProcess()

	// for _, value := range instructions {
	// 	fmt.Println(value)
	// }
}

// func setInstructions() {
// 	instructions = append(instructions, "FROM ubuntu:latest")
// 	instructions = append(instructions, `CMD ["sleep", "infinity"]`)
// }

// func addInstructions() {

// }
