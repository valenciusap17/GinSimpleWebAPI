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
	respData, err := s.HandleCreateUser(*queryData)
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
