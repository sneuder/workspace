package settings

import (
	"log"
	"os"
	"workspace/cmd"
	"workspace/constants"
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
			cmd.GetLsContainerCmd(),
		},
	}

	if err := cmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func SetFolder() {
	homerDir := helpers.GetHomeDir()
	helpers.CreateFolder(homerDir, constants.FOLDER_WKSPACE)
}

func SetJsonFile() {
	wkspaceFolder := helpers.GetWkspaceFolder()
	data := []schemas.Workspace{}
	helpers.CreateJSONFile(data, wkspaceFolder, constants.FOLDER_WKSPACE)
}
