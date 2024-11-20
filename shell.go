package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shell struct {
	currentDir string
}

func (s *Shell) Start() {
	for {
		fmt.Printf("%s> ", s.currentDir)
		input := s.readInput()
		if err := s.executeCommand(input); err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func (s *Shell) readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (s *Shell) executeCommand(input string) error {

	if len(input) == 0 {
		return nil
	}

	args := strings.Split(input, " ")
	switch args[0] {
	case "cd":
		return s.changeDirectory(args[1:])
	case "pwd":
		return s.printWorkingDirectory()
	case "exit":
		return s.exitShell()
	default:
		return s.runExternalCommand(args)
	}
}

func (s *Shell) changeDirectory(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("cd: no directory specified")
	}
	input := strings.Join(args, " ")

	if err := os.Chdir(input); err != nil {
		return err
	}

	dir, _ := os.Getwd()
	s.currentDir = dir

	return nil
}

func (s *Shell) printWorkingDirectory() error {
	fmt.Println(s.currentDir)
	return nil
}

func (s *Shell) exitShell() error {
	os.Exit(0)
	return nil
}

func (s *Shell) runExternalCommand(args []string) error {
	fmt.Println(args)
	return nil
}
