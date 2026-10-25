package historycmd

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/store"
	"github.com/mikepepping/daily-goggles/termtables"
)

type History struct{ config cmds.Config }

func New(config cmds.Config) cmds.Command { return History{config} }
func (c History) Execute(_ []string) error {
	if c.config.StoreFilename == "" {
		return errors.New("empty store filename")
	}
	db, err := cmds.Open(c.config)
	if err != nil {
		return err
	}
	defer db.Close()
	rows, err := store.NewHistoryRepo(db).List(context.Background())
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		fmt.Println("No tasks")
		return nil
	}
	table := termtables.Table{}
	table.Resize(4)
	table.SetHeadings([]string{"ID", "NAME", "STATUS", "DONE AT"})
	for i, task := range rows {
		done := ""
		if !task.DoneAt.IsZero() {
			done = task.DoneAt.String()
		}
		table.AddRow([]string{strconv.Itoa(i), task.Name, string(task.State), done})
	}
	table.Print()
	return nil
}
