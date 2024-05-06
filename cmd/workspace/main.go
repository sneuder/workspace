package main

import (
	"fmt"
	"workspace/config"
	"workspace/file"
)

func main() {
	config.Config()
	a := file.ReadWorkspaceInfo([]string{"."})
	fmt.Println(a.Name)
}
