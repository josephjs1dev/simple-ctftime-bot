package lineservice

import (
	"strings"

	"github.com/josephsalimin/simple-ctftime-bot/internal/config"
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	linecmd "github.com/josephsalimin/simple-ctftime-bot/internal/line/command"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/ioc"
)

// ImplService is our implementation of line service
type ImplService struct {
	bot    domain.LineBotClient
	config *config.Config
}

// BuildService creates ImplService that implement domain.LineService
func BuildService(container *ioc.Container) domain.LineService {
	config := container.Get((*config.Config)(nil)).(*config.Config)
	bot := container.Get((*domain.LineBotClient)(nil)).(domain.LineBotClient)

	return &ImplService{bot, config}
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

func (s *ImplService) findCommand(messageCtx *domain.LineTextMessageContext) linecmd.Command {
	if verifyCommandMessage(messageCtx.Message.Text) {
		commmand, parameter := parseMessageToCommandNameAndData(messageCtx.Message.Text)

		cmdBuilder, ok := linecmd.CommandMap[commmand]
		if !ok {
			return nil
		}

		return cmdBuilder(parameter)
	}

	return nil
}
