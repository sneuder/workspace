package app

import (
	"workspace/inputs"
)

func StartApp() {
	collectionInputs := inputs.SetInputs()
	listenInputUser(collectionInputs)
}
