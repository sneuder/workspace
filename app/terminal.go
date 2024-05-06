package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"workspace/models"
)

func listenInputUser(collectionInputs models.CollectionInputs) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("workspace > ")

		input := receiveInput(reader)
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		inputs := strings.Split(input, " ")
		inputAction, exists := collectionInputs[inputs[0]]

		if !exists {
			fmt.Println("command not found")
			continue
		}

		inputAction.Run(inputs)
	}
}

func receiveInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')

	if err != nil {
		println("error reading input:", err)
		return "error reading input"
	}

	return input
}
