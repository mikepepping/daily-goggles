package cmds

type CmdConfig struct {
	StorePath string
	StoreFilename string
}

type Command interface {
	Execute(args []string)
}

type BuildFunc func(CmdConfig) Command

