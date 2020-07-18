package linecmd

// CommandBuilder is builder function to generate command processor
type CommandBuilder func([]string) Command

// CommandMap maps all command name to command builder
var CommandMap map[string]CommandBuilder = map[string]CommandBuilder{
	"current_event": buildCommandEvent,
}
