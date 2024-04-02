package tasks

import (
	"encoding/json"
	"errors"
	"time"
)

type TaskState string

const (
	Todo TaskState = "todo"
	Done TaskState = "done"
)

type Task struct {
	Name   string    `json:"name"`
	State  TaskState `json:"state"`
	DoneAt time.Time `json:"doneAt"`
}

func (t Task) Stringify() (string, error) {
	str, err := json.Marshal(t)
	if err != nil {
		return "", errors.New("failed to stringify struct")
	}

	return string(str), nil
}
