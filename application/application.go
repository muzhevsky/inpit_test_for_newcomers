package main

import (
	"MyServer/internal/serverMain"
	"fmt"
	"github.com/BurntSushi/toml"
)

func main() {
	err := runApp()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func runApp() error {
	server := serverMain.Server{}
	serverDto := serverMain.Dto{}
	toml.DecodeFile("config/serverConfig.toml", &serverDto)
	server.Initialize(&serverDto)
	err := server.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
