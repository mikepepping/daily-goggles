package cmds

import (
	"fmt"
	"os"
	"strings"
)

type PrintCmd struct{
	config CmdConfig
}

func BuildPrintCmd(config CmdConfig) Command {
	return PrintCmd{config}
}

func (pc PrintCmd) Execute(args []string) {
	if len(args) == 0 {
		fmt.Println("Missing string to print")
		os.Exit(1)
	}
	fmt.Println(strings.Join(args, " "))
}

