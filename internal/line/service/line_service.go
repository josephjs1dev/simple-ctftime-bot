package lineservice

import (
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

func (s *ImplService) findCommand(messageCtx *TextMessageContext) linecmd.Command {
	return nil
}
