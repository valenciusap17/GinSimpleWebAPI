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
	http.HandleFunc("/createUser", s.HandlePostUser)
	http.HandleFunc("/getAllUser", s.HandleGetAllUSer)
	http.HandleFunc("/getUserById", s.HandleGetUSerById)
	http.HandleFunc("/updateUser", s.HandleUpdateUser)
	http.HandleFunc("/deleteUser", s.HandleDeleteUser)
	return http.ListenAndServe(s.listenAddres, nil)
}
