package content

import (
	"simple-ctftime-bot/app/config"
	"simple-ctftime-bot/app/line"
)

// AppContent is used to as our application's content that has information about config, etc. that's used by many functionalities
type AppContent struct {
	ConfigService *config.Service
	Config        config.Config
	Line          line.BotClient
}

func initializeConfig() (*config.Service, error) {
	return config.ReadConfig(&config.EnvReader{})
}

func initializeLineBot(config config.Config) (line.BotClient, error) {
	return line.InitializeBot(config)
}

// InitializeAppContent runs initialization of our application's context
func InitializeAppContent() (*AppContent, error) {
	configService, err := initializeConfig()
	if err != nil {
		return nil, err
	}

	config := configService.GetConfig()
	line, err := initializeLineBot(*config)

	return &AppContent{
		ConfigService: configService,
		Config:        *config,
		Line:          line,
	}, nil
}
