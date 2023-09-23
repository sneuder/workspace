package cmd

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"workspace/internal/cmd/action"
)

var actionsMap = make(map[string]Action)

func StartTerminal() {
	setSetting()

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

type Action func()

func receiveAction(actionKey string) {
  // if action, exists := actionsMap[input]; exists {
  //   action()
  // } else {
  //   fmt.Println("command not found")
  // }

	action, exists := actionsMap[actionKey];

	if !exists {
		fmt.Println("command not found")
		return
	}
	
	action()
}

func setSetting() {
	setActions()
}

func setActions() {
	// basics
	actionsMap["version"] = action.Version
	actionsMap["help"] = action.Help
	actionsMap["clear"] = action.Clear
	actionsMap["exit"] = action.Exit

	// workspace
	actionsMap["workspace"] = action.Workspace
}
