package main

import (
	"workspace/helpers"
	"workspace/settings"
)

func main() {
	settings.SetFolder()

	wkspaceFolder := helpers.GetWkspaceFolder()
	if checkedJsonFile := helpers.CheckJSONFile(wkspaceFolder, "wkspace"); !checkedJsonFile {
		settings.SetJsonFile()
		return
	}

	settings.SetCmd()
}
