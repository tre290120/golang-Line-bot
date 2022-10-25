package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoClient *mongo.Client

func initConnection() {
	var err error
	client := options.Client().ApplyURI("mongodb://localhost:27017")

	mgoClient, err = mongo.Connect(context.TODO(), client)
	if err != nil {
		log.Fatal(err)
	}

	err = mgoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getMgoCli() *mongo.Client {
	if mgoClient == nil {
		initConnection()
	}
	return mgoClient
}
