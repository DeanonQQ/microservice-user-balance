package apiserver

import (
	"net/http"

	"github.com/deanonqq/microservice-user-balance/config"
)

type APIServer struct {
	config *config.Config
}

func New(config *config.Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()
	return http.ListenAndServe(s.config.Port, nil)
}

func (s *APIServer) configureRouter() {
	http.HandleFunc("/report/users/", s.handleUser())
}
