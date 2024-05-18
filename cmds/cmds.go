package cmds

type Config struct {
	StorePath     string
	StoreFilename string
}

type Command interface {
	Execute(args []string) error
}

type BuildFunc func(Config) Command
