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
	input := strings.Join(args, " ")

	if err := os.Chdir(input); err != nil {
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

	fmt.Println(currentDir)
	return NoOutput, nil
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
	for _, dir := range dirList {
		if !showHidden && strings.HasPrefix(dir.Name(), ".") {
			continue
		}

		if dir.IsDir() {
			fmt.Println(dir.Name() + "/")
		} else {
			fmt.Println(dir.Name())
		}

	}
	return NoOutput, nil
}

func cmdExe(args []string) (string, error) {
	cmd := exec.Command(args[0], args[1:]...)
	stdoutStderr, err := cmd.CombinedOutput()
	return string(stdoutStderr), err

}
