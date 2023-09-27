package docker

import (
	"fmt"
	"log"
	"os/exec"
	"workspace/internal/file"
	"workspace/internal/model"
)

func StartImageProcess(dataWorkspace map[string]model.DataWorkspace) {
	file.Open(dataWorkspace["name"].Value)

	setImage(dataWorkspace["image"].Value)
	setUpdate()
	setTools(dataWorkspace["tools"].Value)

	setCMD()

	file.Close()

	// buildImage()
}

func setImage(value string) {
	image := "FROM " + value
	file.Write([]byte(image))
}

func setUpdate() {
	updata := "RUN apt-get update"
	file.Write([]byte(updata))
}

func setTools(tools string) {
	toolToInstall := "RUN apt install " + tools + " -y"
	println(toolToInstall)
	file.Write([]byte(toolToInstall))
}

func setCMD() {
	cmd := `CMD ["sleep", "infinity"]`
	file.Write([]byte(cmd))
}

func buildImage() {
	cmd := exec.Command("docker", "build", "-t", "testing", ".")

	output, err := cmd.Output()

	outputStr := string(output)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(outputStr)
}
