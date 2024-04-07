package api

import (
	"github.com/google/uuid"
	"github.com/valenciusap17/GoT_Auth/models"
)

func (s *Server) CreateUser(data models.CreateUserRequest) (*models.User, error) {
	newUser := models.NewUser(data.Username, data.Password, data.FirstName, data.LastName)
	stored, err := s.store.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return stored, nil
}

func (s *Server) GetAllUser() ([]*models.User, error) {
	response, err := s.store.ReadAllUser()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Server) GetUserById(id *uuid.UUID) (*models.User, error) {
	response, err := s.store.ReadUserById(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Server) UpdateUser(data *models.CreateUserRequest, id *uuid.UUID) (*models.User, error) {
	response, err := s.store.UpdateUser(data, id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Server) DeleteUser(id *uuid.UUID) (*models.User, error) {
	response, err := s.store.DeleteUser(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}
