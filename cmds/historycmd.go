package cmds

import (
	"errors"
	"path/filepath"

	tasks "github.com/mikepepping/daily-goggles/tasks"
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

	PrintTable(tf.History)

	return nil
}
