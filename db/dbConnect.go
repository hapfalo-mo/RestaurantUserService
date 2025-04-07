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

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
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
	return MongoClient.Database("steakhouseuser").Collection(name)
}
