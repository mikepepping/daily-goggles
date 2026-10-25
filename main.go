package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mikepepping/daily-goggles/cmds"
	begin "github.com/mikepepping/daily-goggles/cmds/begin"
	clean "github.com/mikepepping/daily-goggles/cmds/clean"
	complete "github.com/mikepepping/daily-goggles/cmds/complete"
	delete "github.com/mikepepping/daily-goggles/cmds/delete"
	history "github.com/mikepepping/daily-goggles/cmds/history"
	insert "github.com/mikepepping/daily-goggles/cmds/insert"
	print "github.com/mikepepping/daily-goggles/cmds/print"
	publish "github.com/mikepepping/daily-goggles/cmds/publish"
	wipe "github.com/mikepepping/daily-goggles/cmds/wipe"
	"github.com/mikepepping/daily-goggles/store"
)

func getCmd(name string, resources cmds.Resources) (cmds.Command, error) {
	availableCmds := map[string]cmds.BuildFunc{
		"list":     print.New,
		"ls":       print.New,
		"print":    print.New,
		"add":      insert.New,
		"insert":   insert.New,
		"complete": complete.New,
		"begin":    begin.New,
		"start":    begin.New,
		"publish":  publish.New,
		"delete":   delete.New,
		"remove":   delete.New,
		"clean":    clean.New,
		"history":  history.New,
		"wipe":     wipe.New,
	}

	buildFunc, exists := availableCmds[name]
	if !exists {
		return nil, errors.New("command not found")
	}

	return buildFunc(resources), nil
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

	db, err := store.Open(filepath.Join(homeDir, ".daily-goggles", "store.db"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()
	resources := cmds.Resources{
		Tasks:   store.NewTaskRepo(db),
		History: store.NewHistoryRepo(db),
		DB:      db,
	}

	cmd, err := getCmd(cmdArgs[0], resources)
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
