package Complete

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/store"
)

type Complete struct{ config cmds.Config }

func New(config cmds.Config) cmds.Command { return Complete{config} }
func (c Complete) Execute(args []string) error {
	if c.config.StoreFilename == "" {
		return errors.New("empty store filename")
	}
	if len(args) == 0 {
		return errors.New("missing task index")
	}
	i, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("index argument not an integer")
	}
	db, err := cmds.Open(c.config)
	if err != nil {
		return err
	}
	defer db.Close()
	repo := store.NewTaskRepo(db)
	task, err := repo.Get(context.Background(), i)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no task at this index")
		return nil
	}
	if err != nil {
		return err
	}
	task.State = "done"
	task.DoneAt = time.Now()
	if err := repo.Update(context.Background(), i, task); err != nil {
		return err
	}
	fmt.Println("Marked " + task.Name + " as done.")
	return nil
}
