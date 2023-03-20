package database

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DefaultDatabaseRepository struct {
	client *mongo.Client
}

func (repository *DefaultDatabaseRepository) Init(client *mongo.Client) {
	repository.client = client
}

func (repository *DefaultDatabaseRepository) GetQuestions() []byte {
	usersCollection := repository.client.Database("questionsForInpit").Collection("questions")
	cursor, err := usersCollection.Find(context.TODO(), bson.D{}, nil)
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	json, _ := json.Marshal(results)
	return json
}
