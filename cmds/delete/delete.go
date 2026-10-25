package deletecmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/store"
	"github.com/mikepepping/daily-goggles/tasks"
)

type Delete struct{ resources cmds.Resources }

func New(resources cmds.Resources) cmds.Command { return Delete{resources} }
func (c Delete) Execute(args []string) error {
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
		return tasksRepo.Delete(context.Background(), i)
	})
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no task at this index")
		return nil
	}
	if err != nil {
		return err
	}
	fmt.Println("Deleted: " + task.Name)
	return nil
}
