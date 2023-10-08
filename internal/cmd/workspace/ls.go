package workspace

import (
	"fmt"
	"os"
	"strings"
	"workspace/internal/config"
)

func Ls(args []string) {
	folderToRead := config.PathDirs["workspaces"]
	files, _ := os.ReadDir(folderToRead)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		if !strings.HasSuffix(fileName, "workspace") {
			continue
		}

		fileName = strings.Replace(fileName, "-workspace", "", -1)
		fmt.Println(fileName)
	}

}
