package main

import (
	"fmt"
	"go-shell/shell"
	"os"

	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Failed to enable raw mode:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	currentDir, _ := os.Getwd()
	shell := shell.NewShell(currentDir)
	shell.RegisterCommands()
	shell.Start()
}
