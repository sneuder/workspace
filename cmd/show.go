package cmd

import (
	"fmt"
	"workspace/helpers"

	"github.com/urfave/cli/v2"
)

func GetLsContainerCmd() *cli.Command {
	return &cli.Command{
		Name:    "show",
		Aliases: []string{"s"},
		Usage:   "list all containers from workspce",
		Action:  LsContainerAction,
	}
}

func LsContainerAction(cCtx *cli.Context) error {
	wkspaceData := helpers.GetWorkspaces()

	for _, wkspace := range wkspaceData {
		if wkspace.Default {
			wkspace.Name = "*" + wkspace.Name
		}
		fmt.Println(wkspace.Name)
	}

	return nil
}
