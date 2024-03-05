package main

import (
	"fmt"
	"os"

	"github.com/mikepepping/daily-goggles/cmds"
)


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
