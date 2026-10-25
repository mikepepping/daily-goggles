package clean

import (
	"context"
	"errors"
	"fmt"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/store"
)

type Clean struct{ resources cmds.Resources }

func New(resources cmds.Resources) cmds.Command { return Clean{resources} }
func (c Clean) Execute(_ []string) error {
	err := c.resources.DB.Tx(context.Background(), func(t store.TaskRepo, h store.HistoryRepo) error {
		all, err := t.List(context.Background())
		if err != nil {
			return err
		}
		for _, task := range all {
			if task.State == "done" {
				if err := h.Create(context.Background(), task); err != nil {
					return err
				}
			}
		}
		for i := len(all) - 1; i >= 0; i-- {
			if all[i].State != "done" {
				continue
			}
			if err := t.Delete(context.Background(), i); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return errors.New("failed to clean tasks: " + err.Error())
	}
	fmt.Println("Cleaned Tasks")
	return nil
}
