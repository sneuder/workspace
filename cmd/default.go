package cmd

import (
	"log"
	"workspace/helpers"

	"github.com/urfave/cli/v2"
)

func GetDefaultContainerCmd() *cli.Command {
	return &cli.Command{
		Name:    "default",
		Aliases: []string{"d"},
		Usage:   "set default container for workspace",
		Action:  DefaultContainerAction,
	}
}

func DefaultContainerAction(cCtx *cli.Context) error {
	containerName := cCtx.Args().Get(0)
	container, err := helpers.GetContainerByName(containerName)

	if err != nil {
		log.Fatalln("container does not exists")
	}

	wkspaceFolder := helpers.GetWkspaceFolder()
	wkspaceData := helpers.GetWorkspaces()

	for i := range wkspaceData {
		wkspaceData[i].Default = false

		if wkspaceData[i].Id == container.ID {
			wkspaceData[i].Default = true
		}
	}

	helpers.CreateJSONFile(wkspaceData, wkspaceFolder, "wkspace")
	return nil
}
