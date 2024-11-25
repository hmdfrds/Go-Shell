package shell

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

const NoOutput = ""

type Command interface {
	Execute(args []string) (string, error)
}

type CdCommand struct {
	currentDir *string
}

type CommandFunc func(args []string) (string, error)

func (f CommandFunc) Execute(args []string) (string, error) {
	return f(args)
}

func (c *CdCommand) Execute(args []string) (string, error) {
	if len(args) == 0 {
		return NoOutput, fmt.Errorf("cd: no directory specified")
	}
	if err := os.Chdir(args[0]); err != nil {
		return NoOutput, err
	}
	dir, _ := os.Getwd()
	*c.currentDir = dir
	return NoOutput, nil
}

func cmdPwd(args []string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return NoOutput, err
	}
	return currentDir, nil
}

func cmdExit(args []string) (string, error) {
	os.Exit(0)
	return NoOutput, nil
}

func cmdLs(args []string) (string, error) {
	currentDir, err := os.Getwd()

	if err != nil {
		return NoOutput, err
	}

	dirList, err := os.ReadDir(currentDir)
	if err != nil {
		return NoOutput, err
	}

	showHidden := slices.Contains(args, "-a")
	var result strings.Builder
	for _, dir := range dirList {
		if !showHidden && strings.HasPrefix(dir.Name(), ".") {
			continue
		}
		if dir.IsDir() {
			result.WriteString(dir.Name() + "/\n")
		} else {
			result.WriteString(dir.Name() + "\n")
		}
	}
	return result.String(), nil
}

func cmdExe(args []string) (string, error) {
	cmd := exec.Command(args[0], args[1:]...)
	stdoutStderr, err := cmd.CombinedOutput()
	return string(stdoutStderr), err
}

func cmdEcho(args []string) (string, error) {
	return strings.Join(args, " "), nil
}
