package workspace

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

func Create() {
	println("Creating workspace")
	getDataWorkspace()
	docker.StartImageProcess(dataWorkspace)
	// docker.StartContainerProcess()
}

func Run() {

}

func Ls() {

}

func Stope() {

}

func Remove() {

}

func DecribeCMD() {
	println("usage: workspace")
	println("  create 	- Create a workspace.")
	println("  run 			- Run a workspace")
	println("  stope    - Stope a workspace")
	println("  ls    		- List all workspaces")
	println("  remove   - Remove a workspace.")
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
