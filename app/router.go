package main

import (
	"simple-ctftime-bot/app/context"

	"github.com/gorilla/mux"
)

// GetApplicationRouter create application router with its handler and return the router
func GetApplicationRouter(context *context.AppContext) *mux.Router {
	r := mux.NewRouter()

	return r
}
