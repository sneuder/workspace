package action

import (
	"bufio"
	"os"
	"strings"
	"workspace/internal/docker"
	"workspace/internal/model"
)

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

	for key, data := range dataWorkspace {
		print(data.Text)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		data.Value = input
		dataWorkspace[key] = data
	}
}
