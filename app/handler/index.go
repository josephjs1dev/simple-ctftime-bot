package handler

import (
	"net/http"
	"simple-ctftime-bot/app/content"
)

// IndexHandler will handle all request match the URL
type IndexHandler struct {
}

// BuildIndexHandler creates IndexHandler struct
func BuildIndexHandler(content *content.AppContent) *IndexHandler {
	return &IndexHandler{}
}

// Welcome is index page
func (h IndexHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}
