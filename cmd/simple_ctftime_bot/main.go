package main

import (
	"log"

	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/logger"
)

func main() {
	if err := logger.Init(); err != nil {
		log.Fatalf("Failed to initialize logger, error: %v", err)
	}

	defer logger.Sync()

	server, err := CreateServer()
	if err != nil {
		log.Fatalf("Failed to create server, error: %v", err)
	}

	log.Fatalf(server.Run().Error())
}
