package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"workspace/internal/cmd/action"
)

type ActionCMD func()

var actionsCMD = map[string]ActionCMD{
	"workspace": action.Workspace,
	"clear":     action.Clear,
	"version":   action.Version,
	"help":      action.Help,
	"exit":      action.Exit,
}

func StartTerminal() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("workspace > ")

		input := receiveInput(reader)
		input = strings.TrimSpace(input)

		receiveAction(input)
	}
}

func receiveInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("error reading input:", err)
		return "error reading input"
	}

	return input
}

func receiveAction(actionKey string) {
	action, exists := actionsCMD[actionKey]

	if !exists {
		fmt.Println("command not found")
		return
	}

	action()
}
