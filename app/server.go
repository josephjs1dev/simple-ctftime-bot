package main

import (
	"log"
	"simple-ctftime-bot/app/handler"
	"simple-ctftime-bot/app/ioc"
	lineservice "simple-ctftime-bot/app/line/service"

	"github.com/gorilla/mux"
)

func addServices(container *ioc.Container) {
	// Build Services
	lineService := lineservice.BuildService(container)

	// Bind Services
	if err := container.BindInterface(lineService, (*lineservice.Service)(nil)); err != nil {
		log.Fatal(err)
	}
}

func createServer(container *ioc.Container) *mux.Router {
	addServices(container)

	// Build Handler
	indexHandler := handler.BuildIndexHandler()
	lineHandler := handler.BuildLineBotHandler(container)

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler.Welcome())
	r.HandleFunc("/line", lineHandler.Callback()).Methods("POST")

	return r
}
