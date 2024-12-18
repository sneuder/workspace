package helpers

import (
	"log"
	"workspace/constants"
	"workspace/schemas"
)

func GetDefaultWkspace() schemas.Workspace {
	wkspaces := GetWorkspaces()
	foundWkspace := schemas.Workspace{}

	for _, wkspace := range wkspaces {
		if wkspace.Default {
			foundWkspace = wkspace
			break
		}
	}

	return foundWkspace
}

func GetWorkspaces() []schemas.Workspace {
	wkspaceFolder := GetWkspaceFolder()
	wkspaceData := &[]schemas.Workspace{}

	if err := ReadJSONFile(wkspaceFolder, constants.JSON_FILE_WKSPACE, wkspaceData); err != nil {
		log.Fatalf("error when reading workspaces: %v", err)
	}

	return *wkspaceData
}

func SaveWorkspaceData(wkspaceData []schemas.Workspace) {
	wkspaceFolder := GetWkspaceFolder()
	CreateJSONFile(wkspaceData, wkspaceFolder, constants.JSON_FILE_WKSPACE)
}
