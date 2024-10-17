package cmd

import "github.com/urfave/cli/v2"

func GetAddContainerCmd() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a container to workspace",
		Action:  AddContainerAction,
	}
}

func AddContainerAction(cCtx *cli.Context) error {
	return nil
}
