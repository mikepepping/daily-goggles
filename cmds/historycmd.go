package cmds

import (
	"errors"
	"path/filepath"
	"strconv"

	tasks "github.com/mikepepping/daily-goggles/tasks"
	"github.com/mikepepping/daily-goggles/termtables"
)

type HistoryCmd struct {
	config CmdConfig
}

func BuildHistoryCmd(config CmdConfig) Command {
	return HistoryCmd{config}
}

func (hc HistoryCmd) Execute(_ []string) error {
	tf := tasks.BuildTaskFile()

	fpath := filepath.Join(hc.config.StorePath, hc.config.StoreFilename)
	if err := tf.LoadFromFile(fpath); err != nil {
		return errors.New("failed to load task file")
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
