package shell

import (
	"fmt"

	"strings"
)

type Shell struct {
	CurrentDir      string
	commandRegistry map[string]Command
}

func (s *Shell) Start() {
	for {
		fmt.Printf("%s> ", s.CurrentDir)
		input := s.readInput()

		output, err := s.executeCommand(input)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		if len(output) > 0 {
			println(output)
		}
	}
}

func (s *Shell) RegisterCommands() {
	s.commandRegistry = map[string]Command{
		"cd":   &CdCommand{currentDir: &s.CurrentDir},
		"pwd":  CommandFunc(cmdPwd),
		"exit": CommandFunc(cmdExit),
		"ls":   CommandFunc(cmdLs),
	}
}

func (s *Shell) executeCommand(input string) (string, error) {

	if len(input) == 0 {
		return NoOutput, nil
	}
	args := strings.Split(input, " ")
	if cmd, exists := s.commandRegistry[args[0]]; exists {
		return cmd.Execute(args[1:])
	}
	return s.runExternalCommand(args)
}

func (s *Shell) runExternalCommand(args []string) (string, error) {
	return cmdExe(args)
}
