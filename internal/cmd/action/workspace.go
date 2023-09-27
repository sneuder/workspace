package action

import (
	"bufio"
	"os"
	"strings"
	"workspace/internal/docker"
	"workspace/internal/model"
)

var orderToGetData = []string{"name", "image", "tools", "bindMount"}

var dataWorkspace = map[string]model.DataWorkspace{
	"name": {
		Text:     "Name for workspace: ",
		Value:    "",
		Required: true,
	},
	"image": {
		Text:     "Image for workspace: ",
		Value:    "",
		Required: true,
	},
	"tools": {
		Text:     "Tools to include: ",
		List:     []string{"git", "make"},
		Value:    "",
		Required: false,
	},
	"bindMount": {
		Text:     "Path for workspace: ",
		Value:    "",
		Required: true,
	},
}

func Workspace() {
	println("Creating workspace")
	getDataWorkspace()
	docker.StartImageProcess(dataWorkspace)
	// docker.StartContainerProcess()
}

func getDataWorkspace() {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(orderToGetData); i++ {
		data := dataWorkspace[orderToGetData[i]]
		print(data.Text)

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
