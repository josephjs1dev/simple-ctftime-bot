package lineservice

import (
	"simple-ctftime-bot/app/config"
	"simple-ctftime-bot/app/content"
	"simple-ctftime-bot/app/line"
	linecontent "simple-ctftime-bot/app/line/content"
)

// Service is our line service interface that defines function that needs to be implemented
type Service interface {
	HandleIncomingMessage(textMessageContent linecontent.TextMessageContent) error
}

// ImplService is our implementation of line service
type ImplService struct {
	bot    line.BotClient
	config config.Config
}

// BuildService creates ImplService
func BuildService(appContent *content.AppContent) Service {
	return &ImplService{bot: appContent.Line, config: appContent.Config}
}
