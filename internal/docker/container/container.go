package container

import (
	"fmt"
	"log"
	"os/exec"
	"path"
	"strings"
	"workspace/internal/config"
)

var buildContainerCMD = []string{}

func StartContainerProcess(dataContainer map[string]string) {
	setInitializer()
	setBindMount(dataContainer["bindMount"])
	setContainerName(dataContainer["name"])

	fmt.Println(buildContainerCMD)
	buildContainer()
}

func buildContainer() {
	cmd := exec.Command(buildContainerCMD[0], buildContainerCMD[1:]...)

	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
}

func setInitializer() {
	buildContainerCMD = append(buildContainerCMD, "docker", "run", "-d")
}

func setContainerName(workspaceName string) {
	buildContainerCMD = append(buildContainerCMD, "--name", workspaceName, workspaceName)
}

func setExposePort() {

}

func setBindMount(pathBindMount string) {
	fullPathBindMount := path.Join(config.BasePath, pathBindMount)
	bindMountPartCMD := `type=bind,source=` + fullPathBindMount + `,target=/workspace`
	buildContainerCMD = append(buildContainerCMD, "--mount", bindMountPartCMD)
}

func Exists(containerName string) bool {
	cmd := exec.Command("docker", "inspect", containerName)
	_, err := cmd.Output()

	if err == nil {
		return true
	}

	if strings.Contains(err.Error(), "No such object") {
		return false
	}

	return true
}
