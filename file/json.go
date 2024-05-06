package file

import (
	"encoding/json"
	"path"
	"workspace/internal/config"
	"workspace/utils"
)

type GlobalConfigWorksapce struct {
	Name      string          `json:"name"`
	Version   string          `json:"version"`
	Workspace ConfigWorkspace `json:"workspaces"`
}

type ConfigWorkspace struct {
	Name      string   `json:"name"`
	Image     string   `json:"image"`
	BindMount string   `json:"bindMount"`
	Networks  []string `json:"networks"`
	Tools     []string `json:"tools"`
	Ports     []string `json:"ports"`
	Cmds      []string `json:"cmds"`
}

func ReadWorkspaceInfo(args []string) GlobalConfigWorksapce {
	workspacePath := getPathWorkSpaceInfo(args)
	dataContent := ReadFile(workspacePath)

	var configWorkspace GlobalConfigWorksapce
	json.Unmarshal([]byte(dataContent), &configWorkspace)

	return configWorkspace
}

func getPathWorkSpaceInfo(args []string) string {
	argPath := utils.JoinPathArgs(args)
	return path.Join(config.BasePath, argPath, "workspace.json")
}
