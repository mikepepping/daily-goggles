package cmds

import (
	"errors"
	"fmt"
	"strings"
)

type PrintCmd struct {
	config CmdConfig
}

func BuildPrintCmd(config CmdConfig) Command {
	return PrintCmd{config}
}

func (pc PrintCmd) Execute(args []string) error {
	if len(args) == 0 {
		fmt.Println("Missing string to print")
		return errors.New("missing string to print")
	}
	fmt.Println(strings.Join(args, " "))
	return nil
}
