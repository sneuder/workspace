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
		Action: cmd.CmdExecuteAction,
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
	data := []schemas.Workspace{}
	helpers.CreateJSONFile(data, wkspaceFolder, "wkspace")
}
