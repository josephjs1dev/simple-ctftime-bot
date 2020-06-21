package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"simple-ctftime-bot/app/context"

	"github.com/gorilla/handlers"
)

func main() {
	context, err := context.InitializeAppContext()
	if err != nil {
		log.Fatal(err)
	}

	config := context.Config
	addr := config.Host + ":" + strconv.Itoa(config.Port)

	r := GetApplicationRouter(context)
	router := handlers.LoggingHandler(os.Stdout, r)

	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
