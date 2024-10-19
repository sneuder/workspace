package helpers

import (
	"log"
	"workspace/constants"
	"workspace/schemas"
)

func GetDefaultWkspace() schemas.Workspace {
	wkspaces := GetWorkspaces()
	return wkspaces[0]
}

func GetWorkspaces() []schemas.Workspace {
	wkspaceFolder := GetWkspaceFolder()
	wkspaceData := &[]schemas.Workspace{}

	if err := ReadJSONFile(wkspaceFolder, constants.JSON_FILE_WKSPACE, wkspaceData); err != nil {
		log.Fatalf("error when reading workspaces: %v", err)
	}

	return *wkspaceData
}
