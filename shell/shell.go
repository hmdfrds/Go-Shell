package shell

import (
	"fmt"
	"os"

	"strings"
)

type Shell struct {
	currentDir      string
	commandRegistry map[string]Command
	history         []string
}

func NewShell(currectDir string) *Shell {
	return &Shell{
		currentDir:      currectDir,
		commandRegistry: make(map[string]Command),
		history:         []string{},
	}
}

func (s *Shell) Start() {
	for {
		fmt.Printf("%s> ", s.currentDir)
		input := s.readInput()
		s.history = append(s.history, input)
		cmd, fileName, appendMode := parseRedirection(input)

		output, err := s.runCommand(cmd)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		if len(output) > 0 {
			if len(fileName) > 0 {
				handleFileRedirection(appendMode, fileName, output)
			} else {
				println(output)
			}
		}
	}
}

func handleFileRedirection(appendMode bool, fileName string, output string) {
	var flag int
	if appendMode {
		flag = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	} else {
		flag = os.O_TRUNC | os.O_CREATE | os.O_WRONLY
	}

	file, err := os.OpenFile(fileName, flag, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	_, err = file.Write([]byte(output))
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	err = file.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
	}
}

func (s *Shell) RegisterCommands() {
	s.commandRegistry = map[string]Command{
		"cd":   &CdCommand{currentDir: &s.currentDir},
		"pwd":  CommandFunc(cmdPwd),
		"exit": CommandFunc(cmdExit),
		"ls":   CommandFunc(cmdLs),
		"echo": CommandFunc(cmdEcho),
	}
}

func (s *Shell) executeCommand(input string) (output string, err error) {

	if len(input) == 0 {
		return NoOutput, nil
	}
	args, err := parseCommand(strings.TrimSpace(input))
	if err != nil {
		return NoOutput, err
	}
	if cmd, exists := s.commandRegistry[args[0]]; exists {
		return cmd.Execute(args[1:])
	}
	return s.runExternalCommand(args)
}

func (s *Shell) runExternalCommand(args []string) (string, error) {
	return cmdExe(args)
}

func (s *Shell) runCommand(cmd string) (string, error) {
	cmds := strings.Split(cmd, " | ")
	var output string
	var err error
	for _, c := range cmds {
		c += " " + output
		output, err = s.executeCommand(c)
		if err != nil {
			return NoOutput, err
		}
	}
	return output, nil
}

func parseRedirection(input string) (cmd string, fileName string, appendMode bool) {
	if strings.Contains(input, " >>") {
		parts := strings.SplitN(input, " >>", 2)
		cmd = strings.TrimSpace(parts[0])
		fileName = strings.TrimSpace(parts[1])
		appendMode = true
	} else if strings.Contains(input, " >") {
		parts := strings.SplitN(input, " >", 2)
		cmd = strings.TrimSpace(parts[0])
		fileName = strings.TrimSpace(parts[1])
		appendMode = false
	} else {
		cmd = input
		fileName = ""
		appendMode = false
	}
	return cmd, fileName, appendMode
}
