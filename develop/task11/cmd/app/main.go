package main

import (
	"fmt"
	"os"
	"task11/interfaces/httpServer"
	"task11/interfaces/httpServer/configs"
)

func main() {
	config := configs.NewConfig()
	err := config.LoadConfigs("configs.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	server, err := httpServer.NewServer(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	server.Start()
}
