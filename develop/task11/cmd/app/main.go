package main

import (
	"fmt"
	"os"
	"task11/infrastructure/inmemory"
	"task11/interfaces/httpserver"
	"task11/interfaces/httpserver/configs"
)

func main() {
	repo, err := inmemory.NewInMemoryRepo()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config := configs.NewConfig()
	err = config.LoadConfigs("configs.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	server, err := httpserver.NewServer(config, repo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	server.Start()
}
