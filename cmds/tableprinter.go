package cmds

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tasks "github.com/mikepepping/daily-goggles/tasks"
)

func PrintTable(tsks []tasks.Task) {
	if len(tsks) == 0 {
		fmt.Println("No Tasks")
		return
	}

	// first find the longest length of each attribute as a string
	iid := 0
	iname := 1
	istate := 2
	idone := 3

	lengths := []int{len("ID"), len("NAME"), len("STATE"), len("DONE AT")}

	for i, tsk := range tsks {
		if nLen := len(strconv.Itoa(i)); nLen > lengths[iid] {
			lengths[iid] = nLen
		}

		if nLen := len(tsk.Name); nLen > lengths[iname] {
			lengths[iname] = nLen
		}

		if nLen := len(tsk.State); nLen > lengths[istate] {
			lengths[istate] = nLen
		}

		if nLen := len(tsk.DoneAt.Format(time.DateTime)); nLen > lengths[idone] {
			lengths[idone] = nLen
		}
	}

	// print table heading
	heading := fmt.Sprintf("| %s | %s | %s | %s |", padRight("ID", lengths[iid], " "), padRight("NAME", lengths[iname], " "), padRight("STATE", lengths[istate], " "), padRight("DONE AT", lengths[idone], " "))
	fmt.Println(strings.Repeat("-", len(heading)))
	fmt.Println(heading)
	fmt.Println(strings.Repeat("-", len(heading)))

	// print table body
	for i, tsk := range tsks {
		id := strconv.Itoa(i)
		id = padRight(id, lengths[iid], " ")
		name := padRight(tsk.Name, lengths[iname], " ")
		state := padRight(string(tsk.State), lengths[istate], " ")
		date := tsk.DoneAt.Format(time.DateTime) // this should always be the same length?

		fmt.Println("|", id, "|", name, "|", state, "|", date, "|")
	}

	// close table
	fmt.Println(strings.Repeat("-", len(heading)))
}

func padRight(str string, maxLength int, pad string) string {
	padLen := maxLength - len(str)
	padding := strings.Repeat(pad, padLen)
	return str + padding
}
