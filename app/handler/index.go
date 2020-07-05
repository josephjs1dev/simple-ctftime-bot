package handler

import (
	"net/http"
)

// IndexHandler will handle all request match the URL
type IndexHandler struct {
}

// BuildIndexHandler creates IndexHandler struct
func BuildIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

// Welcome is index page
func (h IndexHandler) Welcome() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!"))
	})
}
