package linecmd

type Command interface {
}

type BaseCommand struct {
	Parameter []string
}
