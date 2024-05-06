package docker

import (
	"log"
	"os/exec"
	"path"
	"strings"
	"workspace/internal/config"
)

func StartContainerProcess(dataContainer map[string]string) {
	var buildContainerCMD = []string{}

	buildContainerCMD = append(buildContainerCMD, getInitializer()...)
	buildContainerCMD = append(buildContainerCMD, getBindMount(dataContainer["bindmount"])...)
	buildContainerCMD = append(buildContainerCMD, getExposePorts(dataContainer["ports"])...)
	buildContainerCMD = append(buildContainerCMD, getNetworks(dataContainer["networks"])...)
	buildContainerCMD = append(buildContainerCMD, getContainerName(dataContainer["name"])...)

	buildContainer(buildContainerCMD)
}

func StartContainer(workspaceName string) {
	cmd := exec.Command("docker", "start", workspaceName)
	cmd.Output()
}

func StopContainer(workspaceName string) {
	cmd := exec.Command("docker", "container", "stop", workspaceName)
	cmd.Output()
}

func RemoveContainer(workspaceName string) {
	cmd := exec.Command("docker", "container", "rm", workspaceName)
	cmd.Output()
}

// composers

func buildContainer(buildContainerCMD []string) {
	cmd := exec.Command(buildContainerCMD[0], buildContainerCMD[1:]...)
	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
}

func getInitializer() []string {
	return []string{
		"docker", "run", "-d",
	}
}

func getContainerName(workspaceName string) []string {
	return []string{
		"--name", workspaceName, workspaceName,
	}
}

func getExposePorts(exposePorts string) []string {
	var portsCMD []string

	if exposePorts == "" {
		return portsCMD
	}

	collectionPort := strings.Split(exposePorts, " ")

	for _, port := range collectionPort {
		portsCMD = append(portsCMD, "-p", port+":"+port)
	}

	return portsCMD
}

func getBindMount(pathBindMount string) []string {
	fullPathBindMount := path.Join(config.BasePath, pathBindMount)
	bindMountPartCMD := `type=bind,source=` + fullPathBindMount + `,target=/workspace`

	return []string{
		"--mount", bindMountPartCMD,
	}
}

func getNetworks(networks string) []string {
	var networksCMD []string

	if networks == "" {
		return networksCMD
	}

	collectionNetwork := strings.Split(networks, " ")

	for _, network := range collectionNetwork {
		networksCMD = append(networksCMD, "--network="+network)
	}

	return networksCMD
}
