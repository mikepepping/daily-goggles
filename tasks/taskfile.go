package tasks

import (
	"encoding/json"
	"errors"
)

type TaskFile struct {
	Tasks []Task `json:"tasks"`
}

func (tf TaskFile) Stringify() (string, error) {
	str, err := json.Marshal(tf)
	if err != nil {
		return "", errors.New("failed to stringify taskfile")
	}

	return string(str), nil
}
