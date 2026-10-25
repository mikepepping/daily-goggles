package insertcmd

import (
	"context"
	"errors"
	"strings"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/store"
	"github.com/mikepepping/daily-goggles/tasks"
)

type Insert struct{ config cmds.Config }

func New(config cmds.Config) cmds.Command { return Insert{config} }

func (c Insert) Execute(args []string) error {
	if c.config.StoreFilename == "" {
		return errors.New("empty store filename")
	}
	db, err := cmds.Open(c.config)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Tx(context.Background(), func(repo store.TaskRepo, _ store.HistoryRepo) error {
		return repo.Create(context.Background(), tasks.Task{Name: strings.Join(args, " "), State: tasks.Todo})
	})
}
