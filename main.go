package main

import (
	"fmt"
	"os"
	"strings"
)

type Command interface {
	Execute(args []string)
}


type PrintCommand struct {}

func (pc PrintCommand) Execute(args []string) {
	if len(args) == 0 {
		fmt.Println("Missing string to print")
		os.Exit(1)
	}
	fmt.Println(strings.Join(args, " "))
}

func main() {
	cmdArgs := os.Args[1:]
	if len(cmdArgs) == 0 {
		fmt.Println("No Command Given")
		os.Exit(1)
	}

	commands := map[string]Command{
		"print": PrintCommand{},
	}

	cmd := commands[cmdArgs[0]]
	cmd.Execute(cmdArgs[1:])

	os.Exit(0)
}
