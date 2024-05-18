package insertcmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mikepepping/daily-goggles/cmds"
	tasks "github.com/mikepepping/daily-goggles/tasks"
)

type InsertCmd struct {
	config cmds.Config
}

func New(config cmds.Config) cmds.Command {
	return InsertCmd{config}
}

func (ic InsertCmd) Execute(args []string) error {
	if ic.config.StoreFilename == "" {
		return errors.New("empty store filename")
	}

	fpath := filepath.Join(ic.config.StorePath, ic.config.StoreFilename)

	if _, err := os.Stat(fpath); err != nil {
		if err = ic.createTaskFile(); err != nil {
			fmt.Println(err)
			return err
		}
	}

	var tf = tasks.BuildTaskFile()
	if err := tf.LoadFromFile(fpath); err != nil {
		return errors.New("failed to load task file")
	}

	tf.AppendTask(tasks.Task{
		Name:   strings.Join(args, " "),
		State:  tasks.Todo,
		DoneAt: time.Time{},
	})

	if err := tf.SaveToFile(fpath); err != nil {
		return errors.New("failed to save task file")
	}

	return nil
}

func (ic InsertCmd) createTaskFile() error {
	fmt.Println("Creating Task Store")
	fmt.Println("Creating new store directory")
	fmt.Println(ic.config.StorePath)
	if err := os.MkdirAll(ic.config.StorePath, 0755); err != nil {
		return errors.New("failed to create store directory")
	}

	fpath := filepath.Join(ic.config.StorePath, ic.config.StoreFilename)

	tf := tasks.BuildTaskFile()
	if err := tf.SaveToFile(fpath); err != nil {
		return errors.New("failed to save new task file")
	}

	return nil
}
