package main

import (
	"fmt"
	"log"

	repository "github.com/valenciusap17/GoT_Auth/Repository"
	"github.com/valenciusap17/GoT_Auth/api"
)

func main() {
	storage, err := repository.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(storage)

	if err := storage.Init(); err != nil {
		log.Fatal(err)
	}

	listenAddress := "localhost:8080"
	server := api.NewServer(listenAddress, storage)
	fmt.Println("Server is currently running on port: ", listenAddress)
	server.Start()
}
