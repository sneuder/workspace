package cmd

import (
	"log"
	"workspace/helpers"

	"github.com/urfave/cli/v2"
)

func GetRemoveContainerCmd() *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Aliases: []string{"r"},
		Usage:   "remove a container from workspace",
		Action:  RemoveContainerAction,
	}
}

func RemoveContainerAction(cCtx *cli.Context) error {
	containerName := cCtx.Args().Get(0)
	container, err := helpers.GetContainerByName(containerName)

	if err != nil {
		log.Fatalln("container does not exists")
	}

	wkspaceData := helpers.GetWorkspaces()

	for index, wkspace := range wkspaceData {
		if container.ID == wkspace.Id {
			wkspaceData = append(wkspaceData[:index], wkspaceData[index+1:]...)
			break
		}
	}

	helpers.SaveWorkspaceData(wkspaceData)

	return nil
}
