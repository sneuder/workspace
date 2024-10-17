package cmd

import "github.com/urfave/cli/v2"

func GetDefaultContainerCmd() *cli.Command {
	return &cli.Command{
		Name:    "default",
		Aliases: []string{"d"},
		Usage:   "set default container for workspace",
		Action:  DefaultContainerAction,
	}
}

func DefaultContainerAction(cCtx *cli.Context) error {
	return nil
}
