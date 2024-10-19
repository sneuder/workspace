package helpers

import (
	"github.com/urfave/cli/v2"
)

func GetAllArgs(cCtx *cli.Context) []string {
	args := []string{}

	for i := 0; i < cCtx.Args().Len(); i++ {
		args = append(args, cCtx.Args().Get(i))
	}

	return args
}
