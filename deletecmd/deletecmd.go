package deletecmd

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/tasks"
)

type DeleteCmd struct {
	config cmds.Config
}

func New(config cmds.Config) cmds.Command {
	return DeleteCmd{config}
}

func (pc DeleteCmd) Execute(args []string) error {
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

	task := tf.Tasks[taskIndex]

	tf.Tasks = append(tf.Tasks[:taskIndex], tf.Tasks[taskIndex+1:]...)

	tf.SaveToFile(fpath)

	fmt.Println("Deleted: " + task.Name)

	return nil
}
