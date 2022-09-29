package internal

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
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
}

func (s *Server) Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The server is ready to serve"))
	}
}
