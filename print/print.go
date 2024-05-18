package print

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/mikepepping/daily-goggles/cmds"
	tasks "github.com/mikepepping/daily-goggles/tasks"
	"github.com/mikepepping/daily-goggles/termtables"
)

type PrintCmd struct {
	config cmds.Config
}

func New(config cmds.Config) cmds.Command {
	return PrintCmd{config}
}

func (pc PrintCmd) Execute(_ []string) error {
	if pc.config.StoreFilename == "" {
		return errors.New("empty store filename")
	}

	tf := tasks.BuildTaskFile()
	fpath := filepath.Join(pc.config.StorePath, pc.config.StoreFilename)

	if _, err := os.Stat(fpath); err != nil {
		if err = tf.SaveToFile(fpath); err != nil {
			fmt.Println(err)
			return err
		}
	}

	if err := tf.LoadFromFile(fpath); err != nil {
		return err
	}

	if len(tf.Tasks) == 0 {
		fmt.Println("No tasks")
		return nil
	}

	table := termtables.Table{}
	table.Resize(4)
	table.SetHeadings([]string{"ID", "NAME", "STATUS", "DONE AT"})
	for i, task := range tf.Tasks {
		doneAt := ""
		if !task.DoneAt.IsZero() {
			doneAt = task.DoneAt.String()
		}

		table.AddRow([]string{strconv.Itoa(i), task.Name, string(task.State), doneAt})
	}
	table.Print()

	return nil
}
