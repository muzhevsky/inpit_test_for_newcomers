package main

import (
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
	serverDto := serverMain.Dto{}
	toml.DecodeFile("config/serverConfig.toml", &serverDto)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	usersCollection := client.Database("questionsForInpit").Collection("questions")
	cursor, err := usersCollection.Find(context.TODO(), bson.D{}, nil)
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	fmt.Println(results[0])

	server.Initialize(&serverDto)
	err = server.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
