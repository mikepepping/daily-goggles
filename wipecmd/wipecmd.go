package wipecmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/tasks"
)

type WipeCmd struct {
	config cmds.Config
}

func New(config cmds.Config) cmds.Command {
	return WipeCmd{config}
}

func (pc WipeCmd) Execute(args []string) error {
	if pc.config.StoreFilename == "" {
		return errors.New("empty store filename")
	}

	fpath := filepath.Join(pc.config.StorePath, pc.config.StoreFilename)
	tf := tasks.BuildTaskFile()
	if err := tf.LoadFromFile(fpath); err != nil {
		return err
	}

	if len(tf.Tasks) == 0 {
		fmt.Println("No tasks")
		return nil
	}

	tf.Tasks = make([]tasks.Task, 0)
	tf.History = make([]tasks.Task, 0)

	tf.SaveToFile(fpath)

	fmt.Println("All Tasks have been wiped including history")

	return nil
}
