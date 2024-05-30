package main

import (
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

func getCmd(name string, config cmds.Config) cmds.Command {
	return map[string]cmds.BuildFunc{
		"print":    printcmd.New,
		"insert":   insertcmd.New,
		"complete": completecmd.New,
		"delete":   deletecmd.New,
		"clean":    cleancmd.New,
		"history":  historycmd.New,
		"wipe":     wipecmd.New,
	}[name](config)
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

	cmd := getCmd(cmdArgs[0], config)
	if err := cmd.Execute(cmdArgs[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
