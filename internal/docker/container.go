package docker

import (
	"log"
	"os/exec"
	"path"
	"strings"
	"workspace/internal/config"
	"workspace/internal/model"
)

var buildContainerCMD = []string{}

func StartContainerProcess(dataWorkspace map[string]model.DataWorkspace) {
	setInitializer()
	setContainerName(dataWorkspace["name"].Value)
	setBindMount(dataWorkspace["bindMount"].Value)

	// buildContainer()
}

func buildContainer() {
	cmd := exec.Command(strings.Join(buildContainerCMD, " "))

	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
}

func setInitializer() {
	buildContainerCMD = append(buildContainerCMD, "docker run -d")
}

func setContainerName(workspaceName string) {
	buildContainerCMD = append(buildContainerCMD, "--name", workspaceName, workspaceName)
}

func setExposePort() {

}

func setBindMount(pathBindMount string) {
	fullPathBindMount := path.Join(config.BasePath, pathBindMount)
	bindMountPartCMD := []string{`"type=bind,source=`, fullPathBindMount, `target=$(/workspace)"`}
	buildContainerCMD = append(buildContainerCMD, strings.Join(bindMountPartCMD, ""))
}
