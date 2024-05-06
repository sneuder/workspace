package inputs

import (
	"fmt"
	"workspace/models"

	"github.com/urfave/cli/v2"
)

func setWorkspaceInputs() {
	setBuildWorkspaceInput()
	setRestoreWorkspaceInput()
	setStopWorkspaceInput()
	setRemoveWorkspaceInput()
	collectionInputs["ls"] = setLsWorkspacesInput()
}

func setBuildWorkspaceInput() {

}

func setRestoreWorkspaceInput() {

}

func setStopWorkspaceInput() {

}

func setRemoveWorkspaceInput() {

}

func setLsWorkspacesInput() models.InputCMD {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(cCtx *cli.Context) error {

			name := "Nefertiti"
			if cCtx.NArg() > 0 {
				name = cCtx.Args().Get(0)
			}

			fmt.Println(name)
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "p",
				Value:    "",
				Usage:    "language for the greeting",
				Required: true,
			},
		},
	}

	return app
}
