package main

import (
	"log"
	"simple-ctftime-bot/app/config"
	"simple-ctftime-bot/app/ioc"
	appline "simple-ctftime-bot/app/line"
)

func initializeConfig() *config.Config {
	configService, err := config.ReadConfig(&config.EnvReader{})
	if err != nil {
		log.Fatal(err)
	}

	return configService.GetConfig()
}

func initializeLineBot(config *config.Config) appline.BotClient {
	client, err := appline.InitializeBot(config)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func initializeAppContainer() *ioc.Container {
	config := initializeConfig()
	line := initializeLineBot(config)

	container := ioc.CreateContainer()
	container.Bind(config)
	if err := container.BindInterface(line, (*appline.BotClient)(nil)); err != nil {
		log.Fatal(err)
	}

	return container
}
