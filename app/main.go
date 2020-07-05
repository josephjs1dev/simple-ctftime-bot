package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"simple-ctftime-bot/app/config"

	"github.com/gorilla/handlers"
)

func main() {
	container := initializeAppContainer()
	r := createServer(container)
	router := handlers.LoggingHandler(os.Stdout, r)

	config := container.Get((*config.Config)(nil)).(*config.Config)
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
