package main

import (
	"fmt"
	"os"

	"github.com/mikepepping/daily-goggles/cmds"
)

var config = cmds.CmdConfig{
	StorePath:     "/home/michael/.daily-goggles",
	StoreFilename: "store.json",
}

func getCmd(name string) cmds.Command {
	buildCmd := map[string]cmds.BuildFunc{
		"print":  cmds.BuildPrintCmd,
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

	cmd := getCmd(cmdArgs[0])
	if err := cmd.Execute(cmdArgs[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
