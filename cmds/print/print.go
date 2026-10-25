package printcmd

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/store"
	"github.com/mikepepping/daily-goggles/tasks"
	"github.com/mikepepping/daily-goggles/termtables"
)

type Print struct{ config cmds.Config }

func New(config cmds.Config) cmds.Command { return Print{config} }

func (c Print) Execute(_ []string) error {
	if c.config.StoreFilename == "" {
		return errors.New("empty store filename")
	}
	db, err := cmds.Open(c.config)
	if err != nil {
		return err
	}
	defer db.Close()
	rows, err := store.NewTaskRepo(db).List(context.Background())
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		fmt.Println("No tasks")
		return nil
	}
	printRows(rows)
	return nil
}

func printRows(rows []tasks.Task) {
	table := termtables.Table{}
	table.Resize(4)
	table.SetHeadings([]string{"ID", "NAME", "STATUS", "DONE AT"})
	for i, task := range rows {
		doneAt := ""
		if !task.DoneAt.IsZero() {
			doneAt = task.DoneAt.String()
		}
		table.AddRow([]string{strconv.Itoa(i), task.Name, string(task.State), doneAt})
	}
	table.Print()
}
