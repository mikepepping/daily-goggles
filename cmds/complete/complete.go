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
	"github.com/mikepepping/daily-goggles/tasks"
)

type Complete struct{ resources cmds.Resources }

func New(resources cmds.Resources) cmds.Command { return Complete{resources} }
func (c Complete) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("missing task index")
	}
	i, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("index argument not an integer")
	}
	var task tasks.Task
	err = c.resources.DB.Tx(context.Background(), func(tasksRepo store.TaskRepo, _ store.HistoryRepo) error {
		var err error
		task, err = tasksRepo.Get(context.Background(), i)
		if err != nil {
			return err
		}
		task.State, task.DoneAt = tasks.Done, time.Now()
		return tasksRepo.Update(context.Background(), i, task)
	})
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no task at this index")
		return nil
	}
	if err != nil {
		return err
	}
	fmt.Println("Marked " + task.Name + " as done.")
	return nil
}
