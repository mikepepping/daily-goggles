package completecmd

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/tasks"
)

type CompleteCmd struct {
	config cmds.Config
}

func New(config cmds.Config) cmds.Command {
	return CompleteCmd{config}
}

func (pc CompleteCmd) Execute(args []string) error {
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

	taskIndex, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("index argument not an integer")
	}

	if taskIndex > len(tf.Tasks) || taskIndex < 0 {
		fmt.Println("no task at this index")
		return nil
	}

	tf.Tasks[taskIndex].State = "done"
	tf.Tasks[taskIndex].DoneAt = time.Now()

	tf.SaveToFile(fpath)

	fmt.Println("Marked " + tf.Tasks[taskIndex].Name + " as done.")

	return nil
}
