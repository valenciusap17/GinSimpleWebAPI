package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/valenciusap17/GoT_Auth/models"
)

func (s *Server) HandlePostUser(w http.ResponseWriter, r *http.Request) {
	queryData := new(models.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(queryData); err != nil {
		log.Fatal(err)
		return
	}
	respData, err := s.CreateUser(*queryData)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(respData)

	jsonData, err := json.Marshal(respData)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Vontent-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func (s *Server) HandleGetAllUSer(w http.ResponseWriter, r *http.Request) {
	response, err := s.GetAllUser()
	if err != nil {
		log.Fatal(err)
		return
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (s *Server) HandleGetUSerById(w http.ResponseWriter, r *http.Request) {
	id := &models.UserIdRequest{}

	if err := json.NewDecoder(r.Body).Decode(id); err != nil {
		log.Fatal("inkah? ", err)
		return
	}

	response, err := s.GetUserById(&id.Id)
	if err != nil {
		log.Fatal(err)
		return
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (s *Server) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	queryData := &models.UpdateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(queryData); err != nil {
		log.Fatal(err)
		return
	}
	response, err := s.UpdateUser(&queryData.UserData, &queryData.Id)
	if err != nil {
		log.Fatal(err)
		return
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (s *Server) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	queryData := new(models.UserIdRequest)

	if err := json.NewDecoder(r.Body).Decode(queryData); err != nil {
		log.Fatal(err)
		return
	}
	response, err := s.DeleteUser(&queryData.Id)
	if err != nil {
		log.Fatal(err)
		return
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
