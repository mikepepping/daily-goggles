package printcmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mikepepping/daily-goggles/cmds"
	"github.com/mikepepping/daily-goggles/tasks"
	"github.com/mikepepping/daily-goggles/termtables"
)

type Print struct{ resources cmds.Resources }

func New(resources cmds.Resources) cmds.Command { return Print{resources} }

func (c Print) Execute(_ []string) error {
	rows, err := c.resources.Tasks.List(context.Background())
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		fmt.Println("No tasks")
		return nil
	}
	printRows(rows)
	return nil
}

func printRows(rows []tasks.Task) {
	table := termtables.Table{}
	table.Resize(4)
	table.SetHeadings([]string{"ID", "NAME", "STATUS", "DONE AT"})
	for i, task := range rows {
		doneAt := ""
		if !task.DoneAt.IsZero() {
			doneAt = task.DoneAt.String()
		}
		table.AddRow([]string{strconv.Itoa(i), task.Name, string(task.State), doneAt})
	}
	table.Print()
}
