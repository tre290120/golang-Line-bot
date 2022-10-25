package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	TimeStamp primitive.DateTime `bson:"timestamp"`
	ID        string             `bson:"id"`
	Msg       string             `bson:"msg"`
}
