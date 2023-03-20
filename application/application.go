package main

import (
	"MyServer/database"
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
	serverParams := serverMain.ServerParams{}
	databaseDto := database.DatabaseParams{}
	toml.DecodeFile("config/serverConfig.toml", &serverParams)
	toml.DecodeFile("config/databaseConfig.toml", &databaseDto)

	databaseClient := database.CreateMongoClient(databaseDto)
	databaseHandler := &database.DatabaseHandler{}
	databaseHandler.InitHandler(databaseClient)

	server.Initialize(&serverParams, databaseHandler)

	err := server.Start()
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
