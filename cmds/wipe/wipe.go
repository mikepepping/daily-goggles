package wipecmd

import (
	"context"
	"fmt"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/store"
)

type Wipe struct{ resources cmds.Resources }

func New(resources cmds.Resources) cmds.Command { return Wipe{resources} }
func (c Wipe) Execute(_ []string) error {
	err := c.resources.DB.Tx(context.Background(), func(tasksRepo store.TaskRepo, historyRepo store.HistoryRepo) error {
		tasks, err := tasksRepo.List(context.Background())
		if err != nil {
			return err
		}
		for i := len(tasks) - 1; i >= 0; i-- {
			if err := tasksRepo.Delete(context.Background(), i); err != nil {
				return err
			}
		}
		history, err := historyRepo.List(context.Background())
		if err != nil {
			return err
		}
		for i := len(history) - 1; i >= 0; i-- {
			if err := historyRepo.Delete(context.Background(), i); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	fmt.Println("All Tasks have been wiped including history")
	return nil
}
