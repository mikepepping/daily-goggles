package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mikepepping/daily-goggles/cleancmd"
	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/completecmd"
	"github.com/mikepepping/daily-goggles/deletecmd"
	"github.com/mikepepping/daily-goggles/historycmd"
	"github.com/mikepepping/daily-goggles/insertcmd"
	"github.com/mikepepping/daily-goggles/printcmd"
	"github.com/mikepepping/daily-goggles/wipecmd"
)

func getCmd(name string, config cmds.Config) (cmds.Command, error) {
	availableCmds := map[string]cmds.BuildFunc{
		"list":     printcmd.New,
		"ls":       printcmd.New,
		"print":    printcmd.New,
		"add":      insertcmd.New,
		"insert":   insertcmd.New,
		"complete": completecmd.New,
		"delete":   deletecmd.New,
		"remove":   deletecmd.New,
		"clean":    cleancmd.New,
		"history":  historycmd.New,
		"wipe":     wipecmd.New,
	}

	buildFunc, exists := availableCmds[name]
	if !exists {
		return nil, errors.New("command not found.")
	}

	return buildFunc((config)), nil
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

	var config = cmds.Config{
		StorePath:     filepath.Join(homeDir, ".daily-goggles"),
		StoreFilename: "store.json",
	}

	cmd, err := getCmd(cmdArgs[0], config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := cmd.Execute(cmdArgs[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
