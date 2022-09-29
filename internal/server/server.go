package server

import (
	"github.com/VladislavFirsov/Publisher-subscriber/internal/database"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	config   *Config
	logger   *logrus.Logger
	router   *mux.Router
	database *database.Database
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.Configurelog(); err != nil {
		return err
	}

	s.logger.Info("Server successfully started")
	s.ConfigRouter()

	if err := s.ConfigureDb(); err != nil {
		return err
	}
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) Configurelog() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *Server) ConfigRouter() {
	s.router.HandleFunc("/", s.Home())
	s.router.HandleFunc("/clients", s.AllClients())
	s.router.HandleFunc("/clients/{id}", s.ClientById())
}

func (s *Server) ConfigureDb() error {
	db := database.NewDb(s.config.Database)
	if err := db.Open(); err != nil {
		return err
	}
	s.database = db
	return nil
}
