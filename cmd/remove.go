package cmd

import "github.com/urfave/cli/v2"

func GetRemoveContainerCmd() *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Aliases: []string{"a"},
		Usage:   "remove a container from workspace",
		Action:  RemoveContainerAction,
	}
}

func RemoveContainerAction(cCtx *cli.Context) error {
	return nil
}
