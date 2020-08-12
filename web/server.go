package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/josephsalimin/simple-ctftime-bot/internal/config"
	httphandler "github.com/josephsalimin/simple-ctftime-bot/internal/http_handler"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/ioc"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/logger"
)

// Server is our server application
type Server struct {
	*mux.Router
	Config    *config.Config
	Container *ioc.Container
}

func (s *Server) bindImplementations() error {
	for _, binder := range binders {
		err := binder(s.Container)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) bindRoutes() {
	lineHandler := httphandler.BuildLineBotHandler(s.Container)

	s.HandleFunc("/line", lineHandler.Callback()).Methods("POST")
	s.HandleFunc("/line", lineHandler.Index()).Methods("GET")
}

func (s *Server) bindConfig() error {
	config, err := config.ReadConfig(&config.EnvReader{})
	if err != nil {
		return err
	}

	s.Config = config
	s.Container.Bind(config)

	return nil
}

// Run executes the server
func (s *Server) Run() error {
	router := handlers.LoggingHandler(os.Stdout, s)

	config := s.Container.Get((*config.Config)(nil)).(*config.Config)
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	logger.Infof("Listening on %v", addr)
	return http.ListenAndServe(addr, router)
}

// CreateServer runs server initialization
func CreateServer() (*Server, error) {
	s := &Server{
		Router:    mux.NewRouter(),
		Container: ioc.CreateContainer(),
	}
	if err := s.bindConfig(); err != nil {
		return nil, err
	}

	if err := s.bindImplementations(); err != nil {
		return nil, err
	}

	s.bindRoutes()

	return s, nil
}
