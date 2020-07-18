package web

import (
	"github.com/gorilla/mux"
	"github.com/josephsalimin/simple-ctftime-bot/internal/config"
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/josephsalimin/simple-ctftime-bot/internal/line"
	linehandler "github.com/josephsalimin/simple-ctftime-bot/internal/line/handler"
	lineservice "github.com/josephsalimin/simple-ctftime-bot/internal/line/service"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/ioc"
)

// Server is our server application
type Server struct {
	*mux.Router
	Config    *config.Config
	Container *ioc.Container
}

func (s *Server) bindServices() error {
	// Build Services
	lineService := lineservice.BuildService(s.Container)

	// Bind Services
	if err := s.Container.BindInterface(lineService, (*domain.LineService)(nil)); err != nil {
		return err
	}

	return nil
}

func (s *Server) addRoutes() {
	lineHandler := linehandler.BuildLineBotHandler(s.Container)

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

func (s *Server) bindLineBot() error {
	client, err := line.InitializeBot(s.Config)
	if err != nil {
		return err
	}

	return s.Container.BindInterface(client, (*domain.LineBotClient)(nil))
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

	if err := s.bindLineBot(); err != nil {
		return nil, err
	}

	if err := s.bindConfig(); err != nil {
		return nil, err
	}

	s.addRoutes()

	return s, nil
}
