package docker

import (
	"fmt"
	"log"
	"os/exec"
)

func StartContainerProcess() {
	buildContainer()
}

func buildContainer() {
	cmd := exec.Command("docker", "run", "-d", "--name", "testing", "testing")

	output, err := cmd.Output()

	outputStr := string(output)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(outputStr)
	// docker run -d -p <host_port>:<container_port> --name <container_name> <image>
}
