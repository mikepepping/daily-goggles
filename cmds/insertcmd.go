package cmds

import (
	"fmt"
)

type InsertCmd struct {
	config CmdConfig
}

func BuildInsertCmd(config CmdConfig) Command{
	return InsertCmd{config}
}

func (ic InsertCmd) Execute(args []string) {
	fmt.Println("Insert Command")
}


