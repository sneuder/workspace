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
	container, err := helpers.GetContainerByName(containerName)

	if err != nil {
		log.Fatalln("container does not exists")
	}

	wkspaceData := helpers.GetWorkspaces()
	wkspace := schemas.Workspace{Name: containerName, Id: container.ID, Default: false}

	if len(wkspaceData) == 0 {
		wkspace.Default = true
	}

	wkspaceData = append(wkspaceData, wkspace)
	helpers.SaveWorkspaceData(wkspaceData)

	return nil
}
