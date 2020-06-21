package main

import (
	"simple-ctftime-bot/app/content"
	"simple-ctftime-bot/app/handler"
	lineservice "simple-ctftime-bot/app/line/service"

	"github.com/gorilla/mux"
)

// GetApplicationRouter create application router with its handler and return the router
func GetApplicationRouter(content *content.AppContent) *mux.Router {
	r := mux.NewRouter()

	// Initiate Service
	lineService := lineservice.BuildService(content)

	// Initiate Handler
	indexHandler := handler.BuildIndexHandler(content)
	lineHandler := handler.BuildLineBotHandler(content, lineService)

	r.HandleFunc("/", indexHandler.Welcome)
	r.HandleFunc("/line", lineHandler.Callback).Methods("POST")

	return r
}
