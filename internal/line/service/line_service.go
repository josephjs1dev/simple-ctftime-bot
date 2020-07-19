package lineservice

import (
	"strings"

	"github.com/josephsalimin/simple-ctftime-bot/internal/config"
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/ioc"
)

// ImplService is our implementation of line service
type ImplService struct {
	config *config.Config
	bot    domain.LineBotClient
	mapper domain.LineCommandMapper
}

// BuildService creates ImplService that implement domain.LineService
func BuildService(container *ioc.Container) domain.LineService {
	config := container.Get((*config.Config)(nil)).(*config.Config)
	bot := container.Get((*domain.LineBotClient)(nil)).(domain.LineBotClient)
	mapper := container.Get((*domain.LineCommandMapper)(nil)).(domain.LineCommandMapper)

	return &ImplService{config, bot, mapper}
}

func verifyCommandMessage(message string) bool {
	return len(message) > 0 && message[0] == '!'
}

func parseMessageToCommandNameAndData(message string) (string, []string) {
	textSlice := strings.Split(message, " ")

	if len(textSlice) > 1 {
		return textSlice[0], textSlice[1:]
	}

	return textSlice[0], []string{}
}

func (s *ImplService) findCommand(messageCtx *domain.LineTextMessageContext) domain.LineCommand {
	if verifyCommandMessage(messageCtx.Message.Text) {
		commmand, parameter := parseMessageToCommandNameAndData(messageCtx.Message.Text)

		cmdBuilder := s.mapper.GetCommand(commmand)
		if cmdBuilder == nil {
			return nil
		}

		return cmdBuilder(parameter)
	}

	return nil
}
