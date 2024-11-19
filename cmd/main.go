package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CD = "cd"
)

func main() {

	fmt.Print("Your Input: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	cmd, _ := extractCmdArgs(input)

	fmt.Println("Your command is ", cmd)
}

func extractCmdArgs(input string) (string, []string) {
	if isEmptyString(input) {
		return "", nil
	}

	cmd := strings.Split(input, " ")[0]

	if !contains(getAvailableCommand(), cmd) {
		return "", nil
	}

	return cmd, nil
}

func contains(slice []string, cmd string) bool {
	for _, availableCmd := range slice {
		if availableCmd == cmd {
			return true
		}
	}
	return false
}

func isEmptyString(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func getAvailableCommand() []string {
	return []string{CD}
}
