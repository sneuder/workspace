package settings

import (
	"log"
	"os"
	"workspace/cmd"
	"workspace/helpers"
	"workspace/schemas"

	"github.com/urfave/cli/v2"
)

func SetCmd() {
	cmd := &cli.App{
		Commands: []*cli.Command{
			cmd.GetAddContainerCmd(),
			cmd.GetRemoveContainerCmd(),
			cmd.GetDefaultContainerCmd(),
		},
	}

	if err := cmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func SetFolder() {
	homerDir := helpers.GetHomeDir()
	helpers.CreateFolder(homerDir, ".wkspace")
}

func SetJsonFile() {
	wkspaceFolder := helpers.GetWkspaceFolder()
	data := []schemas.Workspaces{}
	helpers.CreateJSONFile(data, wkspaceFolder, "wkspace")
}
