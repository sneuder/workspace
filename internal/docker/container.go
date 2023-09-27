package docker

import (
	"log"
	"os/exec"
)

func StartContainerProcess() {
	buildContainer()
}

func buildContainer() {
	cmd := exec.Command("docker", "run", "-d", "--name", "testing", "testing")

	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
}
