package main

import (
	"os"
)

func main() {
	currentDir, _ := os.Getwd()

	shell := Shell{currentDir: currentDir}
	shell.Start()
}
