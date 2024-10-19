package cmd

import (
	"log"
	"workspace/helpers"
	"workspace/schemas"

	"github.com/urfave/cli/v2"
)

func GetAddContainerCmd() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a container to workspace",
		Action:  AddContainerAction,
	}
}

func AddContainerAction(cCtx *cli.Context) error {
	containerName := cCtx.Args().Get(0)

	if containerName == "" {
		log.Fatalln("provide a container name")
	}

	wkspaceFolder := helpers.GetWkspaceFolder()
	wkspaceData := helpers.GetWorkspaces()

	container, err := helpers.GetContainerByName("wkspace")

	if err != nil {
		log.Fatalln("container does not exists")
	}

	wkspaceData = append(wkspaceData, schemas.Workspace{Name: containerName, Id: container.ID})
	helpers.CreateJSONFile(wkspaceData, wkspaceFolder, "wkspace")

	return nil
}
