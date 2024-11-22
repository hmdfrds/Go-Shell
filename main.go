package main

import (
	"go-shell/shell"
	"os"
)

func main() {
	currentDir, _ := os.Getwd()
	shell := shell.Shell{CurrentDir: currentDir}
	shell.Start()
}
