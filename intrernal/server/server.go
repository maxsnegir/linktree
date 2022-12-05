package server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"linktree/cmd/config"
	"net/http"
)

type Server struct {
	cfg    config.Config
	logger *logrus.Logger
	router *mux.Router
}

func (s *Server) Start() error {
	s.logger.Infof("Server is starting on http://%s", s.cfg.Server.Address)
	return http.ListenAndServe(s.cfg.Server.Address, s.router)
}

func (s *Server) Stop() error {
	return nil
}

func NewServer(cfg config.Config, logger *logrus.Logger) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		router: mux.NewRouter(),
	}
}
