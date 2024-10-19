package cmd

import (
	"log"
	"workspace/helpers"

	"github.com/urfave/cli/v2"
)

func CmdExecuteAction(cCtx *cli.Context) error {
	cmds := helpers.GetAllArgs(cCtx)

	if len(cmds) == 0 {
		log.Fatalln("no commands added")
	}

	defaultWkspace := helpers.GetDefaultWkspace()
	if err := helpers.ExecCommandInContainer(defaultWkspace.Id, cmds); err != nil {
		log.Fatalln("error when running command")
	}

	return nil
}
