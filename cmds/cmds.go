package cmds

import "github.com/mikepepping/daily-goggles/store"

type Resources struct {
	Tasks   store.TaskRepo
	History store.HistoryRepo
	DB      *store.DB
}

type Command interface {
	Execute(args []string) error
}

type BuildFunc func(Resources) Command
