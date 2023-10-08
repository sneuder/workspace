package workspace

import (
	"workspace/internal/model"
)

var orderToGetData = []string{"name", "image", "tools", "bindMount"}

var dataWorkspace = map[string]model.DataWorkspace{
	"name": {
		Text:     "Name for workspace: ",
		Value:    "",
		Required: true,
	},
	"image": {
		Text:     "Image for workspace: ",
		Value:    "",
		Required: true,
	},
	"tools": {
		Text:     "Tools to include: ",
		List:     []string{"git", "make"},
		Value:    "",
		Required: false,
	},
	"bindMount": {
		Text:     "Path for workspace: ",
		Value:    "",
		Required: true,
	},
}

var dataContainer = map[string]string{
	"name":      "",
	"bindMount": "",
	"port":      "",
}
