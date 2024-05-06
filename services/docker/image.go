package docker

import (
	"os"
	"os/exec"
	"strings"
	"workspace/config"
	"workspace/file"
	"workspace/models"
)

func BuildDockerFile(dataWorkspace map[string]models.DataWorkspace) {
	_, newFile := file.OpenFile("dockerfile", config.PathDirs["workspaces"])
	defer file.CloseFile(newFile)

	setImage(dataWorkspace["image"].Value, newFile)
	setUpdate(newFile)
	setTools(dataWorkspace["tools"].Value, newFile)
	setWorkDir(newFile)
	setPorts(dataWorkspace["ports"].Value, newFile)
	setCMD(newFile)
}

func CreateImage(imageName string) {
	cmd := exec.Command("docker", "build", "-t", imageName, config.PathDirs["workspaces"])
	cmd.Output()
}

func RemoveImage(workspaceName string) {
	cmd := exec.Command("docker", "image", "rm", workspaceName)
	cmd.Output()
}

func ExistsImage(imageName string) bool {
	cmd := exec.Command("docker", "inspect", imageName)
	_, err := cmd.Output()
	return err != nil
}

//composers

func setImage(value string, file *os.File) {
	image := "FROM " + value
	file.Write([]byte(image))
}

func setUpdate(file *os.File) {
	updata := "RUN apt-get update"
	file.Write([]byte(updata))
}

func setTools(tools string, file *os.File) {
	if tools == "" {
		return
	}

	toolToInstall := "RUN apt install " + tools + " -y"
	file.Write([]byte(toolToInstall))
}

func setWorkDir(file *os.File) {
	workdir := "WORKDIR /workspace"
	file.Write([]byte(workdir))
}

func setPorts(ports string, file *os.File) {
	if ports == "" {
		return
	}

	collectionPort := strings.Split(ports, " ")

	for _, port := range collectionPort {
		portToSet := "EXPOSE " + port
		file.Write([]byte(portToSet))
	}
}

func setCMD(file *os.File) {
	cmd := `CMD ["sleep", "infinity"]`
	file.Write([]byte(cmd))
}
