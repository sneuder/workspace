package docker

import (
	"fmt"
	"os/exec"
)

func CreateNetwork(networkName string) {
	cmd := exec.Command("docker", "network", "create", networkName)
	_, err := cmd.Output()

	if err != nil {
		fmt.Println("Error when creating network")
		return
	}
}

func ExistsNetWork(networkName string) bool {
	cmd := exec.Command("docker", "network", "inspect", networkName)
	_, err := cmd.Output()
	return err == nil
}

func RemoveNetwork(networkName string) {
	cmd := exec.Command("docker", "network", "rm", networkName)
	_, err := cmd.Output()

	if err != nil {
		fmt.Println("Error when removing network")
		return
	}
}
