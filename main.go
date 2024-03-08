package main

import (
	"fmt"
	"os"

	"github.com/mikepepping/daily-goggles/cmds"
)

var config = cmds.CmdConfig{
	"~/.daily-goggles",
	"store.json",
}

func getCmd(name string) cmds.Command {
	buildCmd := map[string]cmds.BuildFunc{
		"print": cmds.BuildPrintCmd,
		"insert": cmds.BuildInsertCmd,
	}[name]

	return buildCmd(config)
}

func main() {
	cmdArgs := os.Args[1:]
	if len(cmdArgs) == 0 {
		fmt.Println("No Command Given")
		os.Exit(1)
	}

	commands := map[string]cmds.Command{
		"print": cmds.PrintCmd{},
	}

	cmd := commands[cmdArgs[0]]
	cmd.Execute(cmdArgs[1:])

	os.Exit(0)
}
