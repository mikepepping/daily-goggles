package wipecmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/mikepepping/daily-goggles/cmds"
)

type Wipe struct{ config cmds.Config }

func New(config cmds.Config) cmds.Command { return Wipe{config} }
func (c Wipe) Execute(_ []string) error {
	if c.config.StoreFilename == "" {
		return errors.New("empty store filename")
	}
	db, err := cmds.Open(c.config)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.ExecContext(context.Background(), "DELETE FROM tasks")
	if err == nil {
		_, err = db.ExecContext(context.Background(), "DELETE FROM history")
	}
	if err != nil {
		return err
	}
	fmt.Println("All Tasks have been wiped including history")
	return nil
}
