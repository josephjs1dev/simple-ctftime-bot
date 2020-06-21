package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"simple-ctftime-bot/app/content"

	"github.com/gorilla/handlers"
)

func main() {
	content, err := content.InitializeAppContent()
	if err != nil {
		log.Fatal(err)
	}

	config := content.Config
	addr := config.Host + ":" + strconv.Itoa(config.Port)

	r := GetApplicationRouter(content)
	router := handlers.LoggingHandler(os.Stdout, r)

	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
