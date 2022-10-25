package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var db *mongo.Database
var collection *mongo.Collection

func init() {
	client = getMgoCli()
	db = client.Database("db")
	collection = db.Collection("msg")
}

func Insert(message Message) bool {
	message.TimeStamp = primitive.NewDateTimeFromTime(time.Now())
	if _, err := collection.InsertOne(context.TODO(), message); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Get(id string) ([]Message, bool) {
	filter := bson.M{"id": id}
	msgs := []Message{}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return msgs, false
	}

	defer func() {
		if err := cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	for cursor.Next(context.TODO()) {
		msg := Message{}
		if cursor.Decode(&msg) != nil {
			fmt.Print(err)
			return msgs, false
		}
		msgs = append(msgs, msg)
	}
	return msgs, true
}
