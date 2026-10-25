package insertcmd

import (
	"context"
	"strings"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/store"
	"github.com/mikepepping/daily-goggles/tasks"
)

type Insert struct{ resources cmds.Resources }

func New(resources cmds.Resources) cmds.Command { return Insert{resources} }

func (c Insert) Execute(args []string) error {
	return c.resources.DB.Tx(context.Background(), func(tasksRepo store.TaskRepo, _ store.HistoryRepo) error {
		return tasksRepo.Create(context.Background(), tasks.Task{Name: strings.Join(args, " "), State: tasks.Todo})
	})
}
