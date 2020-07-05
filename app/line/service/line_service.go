package lineservice

import (
	"simple-ctftime-bot/app/config"
	"simple-ctftime-bot/app/ioc"
	appline "simple-ctftime-bot/app/line"
	linecontext "simple-ctftime-bot/app/line/context"
)

// Service is our line service interface that defines function that needs to be implemented
type Service interface {
	HandleIncomingMessage(linecontext.TextMessageContext) error
}

// ImplService is our implementation of line service
type ImplService struct {
	bot    appline.BotClient
	config *config.Config
}

// BuildService creates ImplService
func BuildService(container *ioc.Container) Service {
	config := container.Get((*config.Config)(nil)).(*config.Config)
	bot := container.Get((*appline.BotClient)(nil)).(appline.BotClient)

	return &ImplService{bot, config}
}
