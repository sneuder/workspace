package docker

import (
	"fmt"
	"log"
	"os/exec"
	"workspace/internal/file"
)

func StartImageProcess() {
	file.Open()

	SetImage()
	SetUpdate()
	SetCMD()

	file.Close()

	BuildImage()
}

func SetImage() {
	image := "FROM ubuntu:latest"
	file.Write([]byte(image))
}

func SetUpdate() {
	image := "RUN apt-get update"
	file.Write([]byte(image))
}

func SetCMD() {
	cmd := `CMD ["sleep", "infinity"]`
	file.Write([]byte(cmd))
}

func BuildImage() {
	cmd := exec.Command("docker", "build", "-t", "testing", ".")

	output, err := cmd.Output()

	outputStr := string(output)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(outputStr)
}
