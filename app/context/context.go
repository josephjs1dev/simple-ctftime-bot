package context

import (
	"simple-ctftime-bot/app/config"
	"simple-ctftime-bot/app/line"
)

// AppContext is used to as our application context that has information about config, etc. that's used by many functionalities
type AppContext struct {
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

// InitializeAppContext runs initialization of our application's context
func InitializeAppContext() (*AppContext, error) {
	configService, err := initializeConfig()
	if err != nil {
		return nil, err
	}

	config := configService.GetConfig()
	line, err := initializeLineBot(*config)

	return &AppContext{
		ConfigService: configService,
		Config:        *config,
		Line:          line,
	}, nil
}
