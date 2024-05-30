package historycmd

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/mikepepping/daily-goggles/cmds"
	tasks "github.com/mikepepping/daily-goggles/tasks"
	"github.com/mikepepping/daily-goggles/termtables"
)

type HistoryCmd struct {
	config cmds.Config
}

func New(config cmds.Config) cmds.Command {
	return HistoryCmd{config}
}

func (hc HistoryCmd) Execute(_ []string) error {
	tf := tasks.BuildTaskFile()

	fpath := filepath.Join(hc.config.StorePath, hc.config.StoreFilename)
	if err := tf.LoadFromFile(fpath); err != nil {
		return errors.New("failed to load task file")
	}

	if len(tf.History) == 0 {
		fmt.Println("No tasks")
		return nil
	}

	table := termtables.Table{}
	table.Resize(4)
	table.SetHeadings([]string{"ID", "NAME", "STATUS", "DONE AT"})
	for i, task := range tf.History {
		doneAt := ""
		if !task.DoneAt.IsZero() {
			doneAt = task.DoneAt.String()
		}

		table.AddRow([]string{strconv.Itoa(i), task.Name, string(task.State), doneAt})
	}
	table.Print()
	return nil
}
