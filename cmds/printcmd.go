package cmds

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	tasks "github.com/mikepepping/daily-goggles/tasks"
)

type PrintCmd struct {
	config CmdConfig
}

func BuildPrintCmd(config CmdConfig) Command {
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

	PrintTable(tf.Tasks)

	return nil
}
