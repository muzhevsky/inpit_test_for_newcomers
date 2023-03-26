package main

import (
	"MyServer/database"
	"MyServer/internal/serverMain"
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
