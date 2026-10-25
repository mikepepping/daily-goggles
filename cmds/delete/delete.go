package deletecmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/store"
)

type Delete struct{ config cmds.Config }

func New(config cmds.Config) cmds.Command { return Delete{config} }
func (c Delete) Execute(args []string) error {
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
	if err := repo.Delete(context.Background(), i); err != nil {
		return err
	}
	fmt.Println("Deleted: " + task.Name)
	return nil
}
