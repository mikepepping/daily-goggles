package tasks

import (
	"encoding/json"
	"errors"
	"os"
)

type TaskFile struct {
	Tasks   []Task `json:"tasks"`
	History []Task `json:"history"`
}

func BuildTaskFile() TaskFile {
	return TaskFile{Tasks: []Task{}}
}

func (tf *TaskFile) AppendTask(task Task) {
	tf.Tasks = append(tf.Tasks, task)
}

// Clean - moves all 'done' tasks and places them in the history
func (tf *TaskFile) Clean() {
	done := []Task{}
	todo := []Task{}

	for _, t := range tf.Tasks {
		if t.State == Done {
			done = append(done, t) // This is probably slow
		} else {
			todo = append(todo, t) //this is probably slow
		}
	}

	tf.Tasks = todo
	tf.History = append(tf.History, done...)
}

func (tf TaskFile) stringify() (string, error) {
	str, err := json.Marshal(tf)
	if err != nil {
		return "", errors.New("failed to stringify taskfile")
	}

	return string(str), nil
}

func (tf *TaskFile) LoadFromFile(fpath string) error {
	data, err := os.ReadFile(fpath)
	if err != nil {
		return errors.New("failed to read task file")
	}

	if err = json.Unmarshal(data, tf); err != nil {
		return errors.New("failed to deserialize task file")
	}

	if tf.Tasks == nil {
		tf.Tasks = []Task{}
	}

	return nil
}

func (tf TaskFile) SaveToFile(fpath string) error {
	jsonstr, err := tf.stringify()
	if err != nil {
		return errors.New("failed to stringify task file")
	}

	file, err := os.Create(fpath)
	if err != nil {
		return errors.New("failed to open task file for writing")
	}
	defer file.Close()

	_, err = file.WriteString(jsonstr)
	if err != nil {
		return errors.New("failed to write to task file")
	}

	return nil
}
