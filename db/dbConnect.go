package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDB() (*mongo.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	clientsOption := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientsOption)
	// Check connection
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Ping Db for sure
	if err := client.Ping(ctx, nil); err != nil {
		log.Print("Ping to mongo db faoiled :%v", err)
		return nil, err
	}
	log.Print("Connected to Mongo")
	MongoClient = client
	return client, nil
}

func GetCollectionUser(name string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatal("MongoClient is nil — did you forget to call ConnectDB?")
	}
	return MongoClient.Database("steakhouseuser").Collection(name)
}
