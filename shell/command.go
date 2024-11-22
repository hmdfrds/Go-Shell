package shell

import (
	"fmt"
	"os"
	"strings"
)

type Command interface {
	Execute(args []string) error
}

type CdCommand struct {
	currentDir *string
}

type CommandFunc func(args []string) error

func (f CommandFunc) Execute(args []string) error {
	return f(args)
}

func (c *CdCommand) Execute(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("cd: no directory specified")
	}
	input := strings.Join(args, " ")

	if err := os.Chdir(input); err != nil {
		return err
	}

	dir, _ := os.Getwd()
	*c.currentDir = dir

	return nil
}

func cmdPwd(args []string) error {

	currentDir, err := os.Getwd()

	if err != nil {
		return err
	}

	fmt.Println(currentDir)
	return nil
}

func cmdExit(args []string) error {
	os.Exit(0)
	return nil
}

func cmdLs(args []string) error {
	currentDir, err := os.Getwd()

	if err != nil {
		return err
	}

	dirList, err := os.ReadDir(currentDir)
	if err != nil {
		return err
	}

	showHidden := argsContain(args, "-a")

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
	return nil
}

func argsContain(args []string, value string) bool {
	for _, arg := range args {
		if arg == value {
			return true
		}
	}
	return false
}
