package cmds

import (
	"os"
	"errors"
	"fmt"
	"path/filepath"
  "strings"
  "strconv"

	tasks "github.com/mikepepping/daily-goggles/tasks"
)

type PrintCmd struct {
	config CmdConfig
}

func BuildPrintCmd(config CmdConfig) Command {
	return PrintCmd{config}
}

func (pc PrintCmd) Execute(_ []string) error {
	if pc.config.StoreFilename == "" {
		return errors.New("empty store filename")
	}

	tf := tasks.BuildTaskFile()
	fpath := filepath.Join(pc.config.StorePath, pc.config.StoreFilename)

	if _, err := os.Stat(fpath); err != nil {
		if err = tf.SaveToFile(fpath); err != nil {
			fmt.Println(err)
			return err
		}
	}

	if err := tf.LoadFromFile(fpath); err != nil {
		return err
	}

  if len(tf.Tasks) == 0 {
    fmt.Println("No tasks")
    return nil
  }

  pc.printTasks(tf.Tasks)


	return nil
}

func (pc PrintCmd) printTasks(tsks []tasks.Task) {
  // first find the longest length of each attribute as a string
  iid :=0
  iname := 1

  lengths := []int{ 0, 0 }

  for i, tsk := range tsks {
    lengths[iid] = len(strconv.Itoa(i))

    if nLen := len(tsk.Name); nLen > lengths[iname] {
      lengths[iname] = nLen
    }
  }

  // print the table with all columns as equal length
  for i, tsk := range tsks {
    id := strconv.Itoa(i)
    id = padRight(id, lengths[iid], " ")
    name := padRight(tsk.Name, lengths[iname], " ")
    fmt.Println("|", id, "|", name, "|")
  }
}

func padRight(str string, maxLength int, pad string) string {
  padLen := maxLength - len(str)
  padding := strings.Repeat(pad, padLen)
  return str + padding
}























