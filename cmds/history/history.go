package historycmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/termtables"
)

type History struct{ resources cmds.Resources }

func New(resources cmds.Resources) cmds.Command { return History{resources} }
func (c History) Execute(_ []string) error {
	rows, err := c.resources.History.List(context.Background())
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		fmt.Println("No tasks")
		return nil
	}
	table := termtables.Table{}
	table.Resize(4)
	table.SetHeadings([]string{"ID", "NAME", "STATUS", "DONE AT"})
	for i, task := range rows {
		done := ""
		if !task.DoneAt.IsZero() {
			done = task.DoneAt.String()
		}
		table.AddRow([]string{strconv.Itoa(i), task.Name, string(task.State), done})
	}
	table.Print()
	return nil
}
