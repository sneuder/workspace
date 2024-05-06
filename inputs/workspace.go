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
		Action: func(*cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	return app
}
