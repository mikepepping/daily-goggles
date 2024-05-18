package cleancmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/mikepepping/daily-goggles/cmds"
	tasks "github.com/mikepepping/daily-goggles/tasks"
)

type CleanCmd struct {
	config cmds.Config
}

func New(config cmds.Config) cmds.Command {
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
