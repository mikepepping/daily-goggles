package cmds

import (
	"path/filepath"

	"github.com/mikepepping/daily-goggles/store"
)

type Config struct {
	StorePath     string
	StoreFilename string
}

type Command interface {
	Execute(args []string) error
}

type BuildFunc func(Config) Command

func Open(c Config) (*store.DB, error) {
	return store.Open(filepath.Join(c.StorePath, c.StoreFilename))
}
