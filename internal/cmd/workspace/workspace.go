package workspace

import (
	"bufio"
	"fmt"
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
}

func Run() {
	docker.StartContainerProcess(dataWorkspace)
}

func Ls() {

}

func Stop() {

}

func Remove() {

}

func DecribeCMD() {
	fmt.Println("usage: workspace")
	fmt.Printf("  %-20s- %s\n", "create", "Create a workspace.")
	fmt.Printf("  %-20s- %s\n", "run", "Run a workspace")
	fmt.Printf("  %-20s- %s\n", "stop", "Stop a workspace")
	fmt.Printf("  %-20s- %s\n", "ls", "List all workspaces")
	fmt.Printf("  %-20s- %s\n", "remove", "Remove a workspace.")
}

func getDataWorkspace() {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(orderToGetData); i++ {
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
