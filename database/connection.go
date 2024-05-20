package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB *mongo.Database
)

const mongoURI string = "mongodb://127.0.0.1:27017"

func ConnectDB() (*mongo.Client, *mongo.Database) {
	clientOption := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("allen")
	return client, DB;
}
