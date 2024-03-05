package cmds

type Command interface {
	Execute(args []string)
}