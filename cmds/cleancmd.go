package cmds

import (
	"errors"
	"fmt"
	"path/filepath"

	tasks "github.com/mikepepping/daily-goggles/tasks"
)

type CleanCmd struct {
	config CmdConfig
}

func BuildCleanCmd(config CmdConfig) Command {
	return CleanCmd{config}
}

func (cc CleanCmd) Execute(_ []string) error {
	tf := tasks.BuildTaskFile()

	fpath := filepath.Join(cc.config.StorePath, cc.config.StoreFilename)
	if err := tf.LoadFromFile(fpath); err != nil {
		return errors.New("failed to load task file")
	}

	tf.Clean()

	if err := tf.SaveToFile(fpath); err != nil {
		return errors.New("failed to save cleaned task file")
	}

	fmt.Println("Cleaned Tasks")
	return nil
}
