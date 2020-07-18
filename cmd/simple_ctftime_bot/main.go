package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/josephsalimin/simple-ctftime-bot/internal/config"
	"github.com/josephsalimin/simple-ctftime-bot/web"
)

func main() {
	server, err := web.CreateServer()
	if err != nil {
		log.Fatalf("Failed to create server, error: %v", err)
	}

	router := handlers.LoggingHandler(os.Stdout, server)

	config := server.Container.Get((*config.Config)(nil)).(*config.Config)
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
