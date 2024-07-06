package main

import (
	"book-api/cmd/api"
	"book-api/config"
	"fmt"
	"log"
)

func main() {
	port := fmt.Sprintf(":%s",config.Envs.ServerPort)
	server := api.NewServer(nil, port)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}