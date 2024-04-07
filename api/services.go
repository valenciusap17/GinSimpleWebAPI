package api

import (
	"github.com/valenciusap17/GoT_Auth/models"
)

func (s *Server) HandleCreateUser(data models.CreateUserRequest) (*models.User, error) {
	newUser := models.NewUser(data.Username, data.Password, data.FirstName, data.LastName)
	stored, err := s.store.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	return stored, nil
}
