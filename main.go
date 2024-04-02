package main

import (
	"fmt"
	"os"
	"path/filepath"

	cmds "github.com/mikepepping/daily-goggles/cmds"
)

func getCmd(name string, config cmds.CmdConfig) cmds.Command {
	buildCmd := map[string]cmds.BuildFunc{
		"print":    cmds.BuildPrintCmd,
		"insert":   cmds.BuildInsertCmd,
		"complete": cmds.BuildCompleteCmd,
		"clean":    cmds.BuildCleanCmd,
		"history":  cmds.BuildHistoryCmd,
	}[name]

	return buildCmd(config)
}

func main() {
	cmdArgs := os.Args[1:]
	if len(cmdArgs) == 0 {
		fmt.Println("No Command Given")
		os.Exit(1)
	}

	var homeDir, err = os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var config = cmds.CmdConfig{
		StorePath:     filepath.Join(homeDir, ".daily-goggles"),
		StoreFilename: "store.json",
	}

	cmd := getCmd(cmdArgs[0], config)
	if err := cmd.Execute(cmdArgs[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
