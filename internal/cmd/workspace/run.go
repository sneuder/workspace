package workspace

import (
	"fmt"
	"path"
	"workspace/internal/config"
	"workspace/internal/docker/container"
	"workspace/internal/file"
	"workspace/internal/util"
)

func Run(args []string) {

	if len(args) == 0 {
		fmt.Println("workspace name needed")
		return
	}

	dataContainer["name"] = args[0]

	contentFile := file.Read(path.Join(config.PathDirs["workspaces"], dataContainer["name"]+"-config"))
	contentFileMap := util.StringToMap(contentFile, "=")

	dataContainer["bindMount"] = contentFileMap["BINDMOUNTPATH"]
	dataContainer["ports"] = contentFileMap["EXPOSEPORTS"]

	container.StartContainerProcess(dataContainer)
}
