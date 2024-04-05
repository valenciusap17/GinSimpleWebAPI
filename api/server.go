package api

import (
	"net/http"

	repository "github.com/valenciusap17/GoT_Auth/Repository"
)

type Server struct {
	listenAddres string
	store        repository.Storage
}

func NewServer(listenAddres string, store repository.Storage) *Server {
	return &Server{
		listenAddres: listenAddres,
		store:        store,
	}
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.listenAddres, nil)
}
