package main

import (
	"simple-ctftime-bot/app/context"
	"simple-ctftime-bot/app/handler"

	"github.com/gorilla/mux"
)

// GetApplicationRouter create application router with its handler and return the router
func GetApplicationRouter(context *context.AppContext) *mux.Router {
	r := mux.NewRouter()

	indexHandler := handler.BuildIndexHandler(context)

	r.HandleFunc("/", indexHandler.Welcome)

	return r
}
