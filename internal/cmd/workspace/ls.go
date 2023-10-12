package workspace

import (
	"fmt"
	"os"
	"strings"
	"workspace/internal/config"
	"workspace/internal/docker"
)

func Ls(args []string) {
	folderToRead := config.PathDirs["workspaces"]
	files, _ := os.ReadDir(folderToRead)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		if !strings.HasSuffix(fileName, "workspace") {
			continue
		}

		fileName = strings.Replace(fileName, "-workspace", "", -1)
		fmt.Println(getState("main-workspace"))
		fmt.Println(fileName)
	}

}

type State string

const (
	Inactive  State = "inactive"
	Instanced State = "instanced"
	Built     State = "built"
	Running   State = "running"
)

func getState(workspaceName string) State {
	containerInfo, _ := docker.GetContainerInfo(workspaceName)

	if containerInfo.ID == "" {
		return Inactive
	}

	if containerInfo.State.Status == "" {
		return Instanced
	}

	if containerInfo.State.Status == "exited" {
		return Built
	}

	if containerInfo.State.Status == "running" {
		return Running
	}

	return Instanced
}
