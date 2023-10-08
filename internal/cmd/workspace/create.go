package workspace

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"workspace/internal/config"
	"workspace/internal/docker/image"
	"workspace/internal/file"
)

func Create(args []string) {
	setArgs(args)
	getDataWorkspace()
	image.StartImageProcess(dataWorkspace)
	setConfigFile()
	resetWorkspaceData()
}

func getDataWorkspace() {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(orderToGetData); i++ {
		if dataWorkspace[orderToGetData[i]].Value != "" {
			continue
		}

		data := dataWorkspace[orderToGetData[i]]
		fmt.Print(data.Text)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if data.Required && input == "" {
			i -= 1
			continue
		}

		data.Value = input
		dataWorkspace[orderToGetData[i]] = data
	}
}

func setConfigFile() {
	fileName := dataWorkspace["name"].Value + "-config"
	file.Open(fileName, config.PathDirs["workspaces"])
	file.Write([]byte("BINDMOUNTPATH=" + dataWorkspace["bindMount"].Value))
	file.Close()
}

func setArgs(args []string) {
	if len(args) == 0 {
		return
	}

	workspaceName := args[0]

	data := dataWorkspace["name"]
	data.Value = workspaceName
	dataWorkspace["name"] = data
}

func resetWorkspaceData() {
	for key, itemData := range dataWorkspace {
		itemData.Value = ""
		dataWorkspace[key] = itemData
	}
}
