package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shell struct {
	currentDir      string
	commandRegistry map[string]Command
}

func (s *Shell) Start() {
	s.RegisterCommands()
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

func (s *Shell) RegisterCommands() {
	s.commandRegistry = map[string]Command{
		"cd":   &CdCommand{currentDir: &s.currentDir},
		"pwd":  CommandFunc(cmdPwd),
		"exit": CommandFunc(cmdExit),
	}
}

func (s *Shell) executeCommand(input string) error {

	if len(input) == 0 {
		return nil
	}
	args := strings.Split(input, " ")
	if cmd, exists := s.commandRegistry[args[0]]; exists {
		return cmd.Execute(args[1:])
	}
	s.runExternalCommand(args)
	return nil
}

func (s *Shell) runExternalCommand(args []string) error {
	fmt.Println(args)
	return nil
}
