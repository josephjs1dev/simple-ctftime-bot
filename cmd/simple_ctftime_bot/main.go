package main

import (
	"log"

	"github.com/josephsalimin/simple-ctftime-bot/web"
)

func main() {
	err := web.InitAppLog()
	if err != nil {
		log.Fatalf("Failed to initialize logger, error: %v", err)
	}

	defer web.SyncAppLog()

	server, err := web.CreateServer()
	if err != nil {
		log.Fatalf("Failed to create server, error: %v", err)
	}

	log.Fatalf(server.Run().Error())
}
