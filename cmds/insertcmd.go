package cmds

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	tasks "github.com/mikepepping/daily-goggles/tasks"
)

type InsertCmd struct {
	config CmdConfig
}

func BuildInsertCmd(config CmdConfig) Command {
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
			return errors.New("failed to create new task file")
		}
	}

	tf, err := ic.readTaskFile()
	if err != nil {
		return errors.New("failed to read task file")
	}

	newTask := tasks.Task{
		Name:   strings.Join(args, " "),
		State:  tasks.Todo,
		DoneAt: time.Time{},
	}

	tf.Tasks = append(tf.Tasks, newTask)
	return ic.writeTaskFile(tf)
}

func (ic InsertCmd) createTaskFile() error {
	fmt.Println("Creating Task Store")
	fmt.Println("Creating new store directory")
	fmt.Println(ic.config.StorePath)
	if err := os.MkdirAll(ic.config.StorePath, 0755); err != nil {
		return errors.New("failed to create store directory")
	}

	fpath := filepath.Join(ic.config.StorePath, ic.config.StoreFilename)

	fmt.Println("Creating new store file")
	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer file.Close()

	template, err := json.Marshal(tasks.TaskFile{Tasks: []tasks.Task{}})
	if err != nil {
		return errors.New("failed to stringify template")
	}

	_, err = file.WriteString(string(template))
	if err != nil {
		return errors.New("failed to save new blank template")
	}

	return nil
}

func (ic InsertCmd) readTaskFile() (tasks.TaskFile, error) {
	fpath := filepath.Join(ic.config.StorePath, ic.config.StoreFilename)
	data, err := os.ReadFile(fpath)
	if err != nil {
		return tasks.TaskFile{}, errors.New("failed to read task file")
	}

	var tf tasks.TaskFile
	if err = json.Unmarshal(data, &tf); err != nil {
		return tasks.TaskFile{}, errors.New("failed to deserialize task file")
	}

	return tf, nil
}

func (ic InsertCmd) writeTaskFile(tf tasks.TaskFile) error {
	jsonstr, err := tf.Stringify()
	if err != nil {
		return errors.New("failed to stringify task file")
	}

	file, err := os.Create(filepath.Join(ic.config.StorePath, ic.config.StorePath))
	if err != nil {
		// TODO : Failing here at the moment
		return errors.New("failed to open task file for writing")
	}

	_, err = file.WriteString(jsonstr)
	if err != nil {
		return errors.New("failed to write to task file")
	}

	return nil
}
