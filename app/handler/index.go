package handler

import (
	"net/http"
	"simple-ctftime-bot/app/context"
)

// IndexHandler will handle all request match the URL
type IndexHandler struct {
	context *context.AppContext
}

// BuildIndexHandler return IndexHandler struct
func BuildIndexHandler(context *context.AppContext) *IndexHandler {
	return &IndexHandler{context: context}
}

// Welcome is index page
func (h IndexHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}
