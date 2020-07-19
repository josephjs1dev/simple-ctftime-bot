package linecmd

import "github.com/josephsalimin/simple-ctftime-bot/internal/domain"

// CommandMap maps all command name to command builder
var commandMap map[string]domain.LineCommandBuilder = map[string]domain.LineCommandBuilder{
	"current_event": buildCommandEvent,
}

// CommandMapper is implementation of domain.LineCommandMapper
type CommandMapper struct{}

// BuildCommandMapper returns new instance of LineCommandMapper implementation
func BuildCommandMapper() domain.LineCommandMapper {
	return &CommandMapper{}
}

// GetCommand returns command builder from command name
func (m *CommandMapper) GetCommand(command string) domain.LineCommandBuilder {
	v, ok := commandMap[command]

	if ok {
		return v
	}

	return nil
}
